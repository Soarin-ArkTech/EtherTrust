package ingame

import (
	"fmt"

	"github.com/Soarin-ArkTech/EtherTrust/exchange"
	"github.com/ethereum/go-ethereum/common"
	"go.minekube.com/brigodier"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/command"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

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
