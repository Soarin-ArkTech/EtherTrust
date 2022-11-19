package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EVMGetter interface {
	ICoinTX
	INonceGetter
	IGasGetter
	IChainGetter
}

type IEVMTX interface {
	IRawTXGetter
	EVMGetter
}

type ITokenTX interface {
	ICoinTX
	IContractGetter
}

type ICoinTX interface {
	IWEIGetter
	IWalletGetter
}

type IWEIGetter interface {
	GetWEI() uint64
}

type IWalletGetter interface {
	GetWallet() common.Address
}

type IContractGetter interface {
	GetContract() string
}

type INonceGetter interface {
	GetNonce() uint64
}

type IGasGetter interface {
	GetGasPrice() *big.Int
	GetGasLimit() uint64
}

type IChainGetter interface {
	GetChainID() *big.Int
}

type IRawTXGetter interface {
	GetRawTX() *types.Transaction
}
