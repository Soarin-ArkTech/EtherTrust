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
	Amount          int64
	Data            []byte
	EVMTX
}

type EVMTX struct {
	Nonce    uint64
	GasPrice *big.Int
	GasLimit uint64
	ChainID  *big.Int
}

type RawEtherTX struct {
	EtherTX
	UnsignedTX *types.Transaction
}

type EtherTXBuilder struct {
	EtherTX
}

func (eth *EtherTXBuilder) SetWallet(wallet string) {
	eth.RecipientWallet = common.HexToAddress(wallet)
}

func (eth *EtherTXBuilder) SetNonce(nonce uint64) {
	eth.Nonce = nonce
}

func (eth *EtherTXBuilder) SetAmount(amnt int64) {
	eth.Amount = amnt
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

// Build our Unsigned Transaction
func (newTX EtherTXBuilder) BuildTX() IEVMTX {
	builtTX := EtherTX{
		newTX.GetWallet(),
		newTX.GetWEI(),
		newTX.GetContractData(),
		EVMTX{
			Nonce:    *SeqNonce,
			GasPrice: newTX.SetGasPrice(),
			GasLimit: newTX.SetGasLimit(),
			ChainID:  newTX.SetChain(),
		},
	}
	IncrementNonce()

	return RawEtherTX{builtTX, types.NewTransaction(builtTX.GetNonce(),
		builtTX.GetWallet(), big.NewInt(builtTX.GetWEI()), builtTX.GetGasLimit(), builtTX.GetGasPrice(), builtTX.GetContractData())}
}
