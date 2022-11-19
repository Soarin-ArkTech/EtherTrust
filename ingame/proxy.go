package ingame

import (
	"go.minekube.com/common/minecraft/component/codec/legacy"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func BirthProxy(proxy *proxy.Proxy) *EtherProx {
	return &EtherProx{
		Proxy: proxy,
		legacyCodec: &legacy.Legacy{
			Char:              legacy.AmpersandChar,
			NoDownsampleColor: true,
		},
	}
}

type EtherProx struct {
	*proxy.Proxy
	legacyCodec *legacy.Legacy
}
