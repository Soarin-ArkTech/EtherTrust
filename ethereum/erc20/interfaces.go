package erc20

import "github.com/Soarin-ArkTech/EtherTrust/ethereum/evm"

type IContractGetter interface {
	GetContract() string
}

type ITokenTX interface {
	evm.ICoinTX
	IContractGetter
}
