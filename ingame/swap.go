package ingame

import (
	"fmt"

	"github.com/Soarin-ArkTech/EtherTrust/ethereum/erc20"
	"github.com/Soarin-ArkTech/EtherTrust/exchange"
	. "go.minekube.com/common/minecraft/component"

	"go.minekube.com/brigodier"
	"go.minekube.com/common/minecraft/color"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func sellGameCurrency() brigodier.LiteralNodeBuilder {
	swap := command.Command(func(c *command.Context) error {
		var wETH erc20.WETH
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
				txhash, ok := erc20.TransferERC20(wETH)
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
