package ingame

import (
	"fmt"

	"github.com/Soarin-ArkTech/EtherTrust/exchange"
	"go.minekube.com/brigodier"
	. "go.minekube.com/common/minecraft/component"
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
				Content: fmt.Sprint("Your exchange account has $", userAcc.GetUSD(), " of ETH in it currently, good for you... Want a fucking reward?"),
			})

			return nil
		})),
	)
}

func hasCmdPerm(ethertrust *EtherProx, perm string) brigodier.RequireFn {
	return command.Requires(func(c *command.RequiresContext) bool {
		return !ethertrust.Config().RequireBuiltinCommandPermissions || c.Source.HasPermission(perm)
	})
}
