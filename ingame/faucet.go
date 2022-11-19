package ingame

import (
	"fmt"

	// . "go.minekube.com/common/minecraft/component"

	"go.minekube.com/brigodier"
	"go.minekube.com/gate/pkg/command"
)

const GetFaucetPerm = "ethertrust.command.getdrip"

// func sendFaucetDrip() brigodier.LiteralNodeBuilder {
// 	drip := command.Command(func(c *command.Context) error {
// 		var wETH ether.WETH
// 		wETH.SetPrice(0.15) // $0.05 ETH

// 		player := c.Source.(proxy.Player)
// 		exchangeUser := exchange.LoadUser(player)
// 		amount := c.Int("amount")
// 		wETH.SetAmount(amount)
// 		wETH.SetWallet(exchangeUser.Wallet)

// 		player.HasPermission()

// 		if !ValidateExchange(playerInv, amount) {
// 			return c.Source.SendMessage(&Text{
// 				Content: "Not enough Dreams in your inventory, you broke bitch!",
// 				S:       Style{Color: color.Gold, Italic: True},
// 			})
// 		} else {
// 			return c.Source.SendMessage(&Text{
// 				Content: "The exchange is posting your transaction to the Ethereum blockchain.",
// 				S:       Style{Color: color.LightPurple, Italic: True},
// 			})

// 			err := BurnCurrency(player, amount)
// 			if err != nil {
// 				c.Source.SendMessage(&Text{
// 					Content: "We were unable to burn your in-game Dreams, contact the owner.",
// 					S:       Style{Color: color.Red, Bold: True},
// 				})
// 			}

// 			go func() {
// 				txhash, ok := ether.TransferERC20(wETH)
// 				if !ok {
// 					c.Source.SendMessage(&Text{
// 						Content: "We were unable to send the ETH to your wallet, returning funds.",
// 						S:       Style{Color: color.Red, Bold: True},
// 					})

// 					GiveCurrency(player, amount)
// 					fmt.Println("Unable to send ETH after exchange request, error: ", err)
// 				} else {
// 					c.Source.SendMessage(&Text{
// 						Content: fmt.Sprintf("You have swapped %v Dreams for $%.2f of ETH", c.Int("amount"), wETH.GetUSD()),
// 						S:       Style{Color: color.Gold, Italic: True},
// 					})

// 					fmt.Println(txhash)
// 				}
// 			}()

// 		}
// 		return nil
// 	})

// 	return brigodier.Literal("ether-faucet").
// 		Then(brigodier.Argument("amount", brigodier.Int).Executes(drip))
// }

func BigDrip(p *EtherProx) brigodier.LiteralNodeBuilder {
	drip := command.Command(func(c *command.Context) error {
		fmt.Println("Hi mom")
		return nil
	})

	return brigodier.Literal("anal").
		Requires(hasCmdPerm(p, GetFaucetPerm)).
		Executes(command.Command(func(c *command.Context) error {
			fmt.Println("Fuckfuwfejwf")
			return nil
		})).
		Then(brigodier.Argument("", brigodier.String).
			Executes(drip))
}
