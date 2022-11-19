package ether

import (
	"github.com/ethereum/go-ethereum/common"
)

type ITokenTX interface {
	IWalletGetter
	IBalanceGetter
	IContractGetter
}

type ICoinTX interface {
	IBalanceGetter
	IWalletGetter
}

type IBalanceGetter interface {
	GetWEI() uint64
}

type IWalletGetter interface {
	GetWallet() common.Address
}

type IContractGetter interface {
	GetContract() string
}
