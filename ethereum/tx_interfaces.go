package ether

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EVMReader interface {
	IRecipientReader
	INonceReader
	IAmountReader
	IGasReader
	IChainReader
}

type IUnsignedTX interface {
	IRawTXReader
	EVMReader
}

type IRecipientReader interface {
	GetRecipient() common.Address
}

type INonceReader interface {
	GetNonce() uint64
}

type IAmountReader interface {
	GetAmount() *big.Int
}

type IGasReader interface {
	GetGasPrice() *big.Int
	GetGasLimit() uint64
}

type IChainReader interface {
	GetChainID() *big.Int
}

type IRawTXReader interface {
	GetRawTX() *types.Transaction
}
