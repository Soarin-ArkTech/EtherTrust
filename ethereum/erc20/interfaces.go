package erc20

import ether "github.com/Soarin-ArkTech/EtherTrust/ethereum"

type IContractGetter interface {
	GetContract() string
}

type ITokenTX interface {
	ether.ICoinTX
	IContractGetter
}
