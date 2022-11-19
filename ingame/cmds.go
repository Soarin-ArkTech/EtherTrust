package ingame

import (
	"fmt"

	ether "github.com/Soarin-ArkTech/EtherTrust/ethereum"
	"github.com/Soarin-ArkTech/EtherTrust/exchange"
	"github.com/ethereum/go-ethereum/common"
	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/common/minecraft/component/codec/legacy"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// Register command across proxy
func (p *EtherProx) grabBalance() {
	// Create the /ethereum exchange command
	p.Command().Register(brigodier.Literal("ether-bal").
		Executes(command.Command(func(c *command.Context) error {
			userAcc := exchange.LoadUser(c.Source.(proxy.Player))

			c.Source.(proxy.Player).SendMessage(&Text{
				Content: fmt.Sprint("Your exchange account has $", userAcc.WalletBalanceUSD(), " of ETH in it currently, good for you... Want a fucking reward?"),
			})

			return nil
		})),
	)
}

func (p *EtherProx) EtherTrust() {

	p.Command().Register(makeExchangeUser())
	p.Command().Register(sellGameCurrency())
	p.Command().Register(BigDrip(p))
	p.grabBalance()
}

func makeExchangeUser() brigodier.LiteralNodeBuilder {
	makeUser := command.Command(func(c *command.Context) error {
		player := c.Source.(proxy.Player)

		Exchange, err := exchange.LoadAccounts()
		if err != nil {
			player.SendMessage(&Text{Content: "We were unable to load the exchange DB to process this request, please contact the owner!"})
		}

		userAcc := exchange.ExchangeAccountBuilder{
			UUID:   player.ID().String(),
			Wallet: common.HexToAddress(c.String("wallet")),
		}
		userAcc.Build(Exchange)

		completedMsg := c.Source.SendMessage(&Text{
			Content: fmt.Sprintf("Your Ethereum wallet has been linked with your user, %s.", player.Username()),
		})

		return completedMsg
	})

	return brigodier.Literal("ether-set").
		Then(brigodier.Argument("wallet", brigodier.String).Executes(makeUser))
}

func sellGameCurrency() brigodier.LiteralNodeBuilder {
	swap := command.Command(func(c *command.Context) error {
		var wETH ether.WETH
		wETH.SetPrice(0.15) // $0.05 ETH

		player := c.Source.(proxy.Player)
		exchangeUser := exchange.LoadUser(player)
		amount := c.Int("amount")
		wETH.SetAmount(amount)
		wETH.SetWallet(exchangeUser.Wallet)
		wETH.SetContract("0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa")

		playerInv, err := GrabInventory(player)
		if err != nil {
			c.Source.SendMessage(&Text{
				Content: "The exchange failed to grab your inventory!",
			})
		}

		if !ValidateExchange(playerInv, amount) {
			return c.Source.SendMessage(&Text{
				Content: "Not enough Dreams in your inventory, you broke bitch!",
				S:       Style{Color: color.Gold, Italic: True},
			})
		} else {
			c.Source.SendMessage(&Text{
				Content: "The exchange is posting your transaction to the Ethereum blockchain.",
				S:       Style{Color: color.LightPurple, Italic: True},
			})

			err = BurnCurrency(player, amount)
			if err != nil {
				c.Source.SendMessage(&Text{
					Content: "We were unable to burn your in-game Dreams, contact the owner.",
					S:       Style{Color: color.Red, Bold: True},
				})
			}

			go func() {
				txhash, ok := ether.TransferERC20(wETH)
				if !ok {
					c.Source.SendMessage(&Text{
						Content: "We were unable to send the ETH to your wallet, returning funds.",
						S:       Style{Color: color.Red, Bold: True},
					})

					GiveCurrency(player, amount)
					fmt.Println("Unable to send ETH after exchange request, error: ", err)
				} else {
					c.Source.SendMessage(&Text{
						Content: fmt.Sprintf("You have swapped %v Dreams for $%.2f of ETH", c.Int("amount"), wETH.GetUSD()),
						S:       Style{Color: color.Gold, Italic: True},
					})

					fmt.Println(txhash)
				}
			}()

		}
		return nil
	})

	return brigodier.Literal("ether-swap").
		Then(brigodier.Argument("amount", brigodier.Int).Executes(swap))
}

func hasCmdPerm(ethertrust *EtherProx, perm string) brigodier.RequireFn {
	return command.Requires(func(c *command.RequiresContext) bool {
		return !ethertrust.Config().RequireBuiltinCommandPermissions || c.Source.HasPermission(perm)
	})
}

type EtherProx struct {
	*proxy.Proxy
	legacyCodec *legacy.Legacy
}
