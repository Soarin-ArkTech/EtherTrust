package ingame

import (
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/component/codec/legacy"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// Register event subscribers
func (p EtherProx) registerSubscribers() error {
	// Show community treasury at top of screens
	event.Subscribe(p.Event(), 0, p.treasuryBar())

	return nil
}

func BirthProxy(proxy *proxy.Proxy) *EtherProx {
	return &EtherProx{
		Proxy: proxy,
		legacyCodec: &legacy.Legacy{
			Char:              legacy.AmpersandChar,
			NoDownsampleColor: true,
		},
	}
}

func (p *EtherProx) Init() error {
	// Initialize our commands and event ubscribers
	p.EtherTrust()
	return p.registerSubscribers()
}
