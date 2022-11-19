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

// func (ether *EtherTXBuilder) GetAmount() uint64 {
// 	ether.Amount = big.NewInt(int64(amnt))
// }

func (ether EtherTX) GetWEI() int64 {
	return ether.Amount
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

func (eth EtherTX) GetContractData() []byte {
	return eth.Data
}

func (ether RawEtherTX) GetRawTX() *types.Transaction {
	return ether.UnsignedTX
}

func (ether RawEtherTX) GetPrivKeyAddress() common.Address {
	return EVMClient.GetPrivKeyAddress()
}
