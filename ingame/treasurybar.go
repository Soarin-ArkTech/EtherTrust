package ingame

import (
	"context"
	"fmt"
	"strconv"
	"time"

	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	ether "github.com/Soarin-ArkTech/ethereal-dreams/ethereum"

	"go.minekube.com/common/minecraft/color"
	. "go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/bossbar"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func (p *EtherProx) treasuryBar() func(*proxy.LoginEvent) {
	updateTreasury := func(bar bossbar.BossBar, player proxy.Player) {
		// treasury := exchange.ExchangeAccount{
		// 	UUID:   "Ethereal-Dreams-Treasury",
		// 	Wallet: common.HexToAddress("0x16Cde118c2ACc7810591687156597f3BfB301193"),
		// }

		ETHPrice, _ := strconv.ParseFloat(*etAPI.Ethereum.Amount, 32)

		// Treasury Exchange Balance
		treasuryETH := etAPI.BigToFloat32(ether.WeiToNorm(etAPI.TreasuryWrappedETH))

		fmt.Printf("Priceo of Ether: %v\n Treasury: %v\n", float32(ETHPrice), treasuryETH)

		text := &Text{Extra: []Component{
			&Text{
				Content: "Ethereal Treasury @ ",
				S:       Style{Color: color.Gold, Bold: True},
			},
			&Text{
				Content: fmt.Sprintf("$%.2f", treasuryETH*float32(ETHPrice)),
				S:       Style{Color: color.DarkGreen, Bold: True},
			},
		}}
		bar.SetName(text)

		treasuryETHSpot := (float32(treasuryETH * float32(ETHPrice) * 0.01))

		if treasuryETHSpot >= 1 {
			bar.SetPercent(1)
		}

		bar.SetPercent(treasuryETHSpot)
	}

	return func(e *proxy.LoginEvent) {
		if !e.Allowed() {
			return
		}
		player := e.Player()

		// Community Treasury Bar
		tBar := bossbar.New(
			&Text{},
			1,
			bossbar.GreenColor,
			bossbar.Notched10Overlay,
		)

		updateTreasury(tBar, player)
		_ = tBar.AddViewer(player)

		go tick(player.Context(), time.Second*5, func() {
			updateTreasury(tBar, player)
		})
	}
}

// tick runs a function every interval until the context is cancelled.
func tick(ctx context.Context, interval time.Duration, fn func()) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn()
		case <-ctx.Done():
			return
		}
	}
}
