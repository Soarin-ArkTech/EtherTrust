package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EtherTX struct {
	RecipientWallet common.Address
	Nonce           uint64
	Amount          *big.Int
	GasPrice        *big.Int
	GasLimit        uint64
	ChainID         *big.Int
	Data            []byte
}

type EtherTXBuilder struct {
	EtherTX
}

type RawEtherTX struct {
	EtherTX
	UnsignedTX *types.Transaction
}

func (ether *EtherTXBuilder) SetWallet(wallet string) {
	ether.RecipientWallet = common.HexToAddress(wallet)
}

func (ether *EtherTXBuilder) SetNonce(nonce uint64) {
	ether.Nonce = nonce
}

func (ether *EtherTXBuilder) SetAmount(amnt uint64) {
	ether.Amount = big.NewInt(int64(amnt))
}

func (eth *EtherTXBuilder) SetGasPrice() *big.Int {
	gas, err := EVMClient.Client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("Unable to suggest a gas price! ", err)
	}

	return gas.SetUint64(gas.Uint64() + 2000000000)
}

func (eth *EtherTXBuilder) SetGasLimit() uint64 {
	gasLimit, err := EVMClient.Client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: EVMClient.PubKey.Address, To: &eth.RecipientWallet, Data: eth.Data})

	if err != nil {
		fmt.Println("Unable to estimate gas limit! Error: ", err)
	}

	fmt.Println("Gas Limit ", gasLimit)

	return gasLimit
}

func (eth *EtherTXBuilder) SetChain() *big.Int {
	Chain, err := EVMClient.Client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("Unable to set the Chain ID. Error: ", err)
	}

	return Chain
}

func (eth *EtherTXBuilder) SetData(data []byte) {
	eth.Data = data
}

func (ethertx EtherTXBuilder) BuildTX() RawEtherTX {
	txStruct := EtherTX{
		RecipientWallet: ethertx.RecipientWallet,
		Nonce:           *SeqNonce,
		Amount:          ethertx.Amount,
		GasPrice:        ethertx.SetGasPrice(),
		GasLimit:        ethertx.SetGasLimit(),
		ChainID:         ethertx.SetChain(),
		Data:            ethertx.Data,
	}
	IncrementNonce()

	fmt.Println(*SeqNonce)

	return RawEtherTX{txStruct, types.NewTransaction(txStruct.GetNonce(),
		txStruct.GetWallet(), big.NewInt(int64(txStruct.GetWEI())), txStruct.GetGasLimit(), txStruct.GetGasPrice(), txStruct.Data)}
}
