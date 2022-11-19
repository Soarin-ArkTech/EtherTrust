package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (ether EtherTX) GetWallet() common.Address {
	return ether.RecipientWallet
}

func (ether EtherTX) GetNonce() uint64 {
	return ether.Nonce
}

func (ether EtherTX) GetWEI() uint64 {
	return ether.Amount.Uint64()
}

func (ether EtherTX) GetGasPrice() *big.Int {
	return ether.GasPrice
}

func (ether EtherTX) GetGasLimit() uint64 {
	return ether.GasLimit
}

func (ether EtherTX) GetChainID() *big.Int {
	return ether.ChainID
}

func (ether RawEtherTX) GetRawTX() *types.Transaction {
	return ether.UnsignedTX
}
