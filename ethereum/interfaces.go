package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Composed Interfaces
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

type ICoinTX interface {
	IWEIGetter
	IWalletGetter
}

// Segmented Interfaces
type IWEIGetter interface {
	GetWEI() uint64
}

type IWalletGetter interface {
	GetWallet() common.Address
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
