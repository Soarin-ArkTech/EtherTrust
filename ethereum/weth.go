package ether

import (
	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	"github.com/ethereum/go-ethereum/common"
)

type WETH TokenExchange

func (t WETH) GetPrice() float32 {
	return t.Price
}

func (t WETH) GetAmount() float32 {
	return t.Amount
}

func (t WETH) GetWEI() uint64 {
	return NormToWei(t.Amount) // WEI Amount to Send
}

func (t WETH) GetWallet() common.Address {
	return t.Wallet
}

func (t WETH) GetContract() string {
	return t.Contract
}

func (t WETH) GetUSD() float32 {
	return t.Amount * etAPI.Ethereum.CBToFloat32()
}

func (t *WETH) SetPrice(price float32) {
	t.Price = price / etAPI.Ethereum.CBToFloat32()
}

func (t *WETH) SetWallet(wallet common.Address) {
	t.Wallet = wallet
}

func (t *WETH) SetAmount(amnt int) {
	t.Amount = float32(amnt) * t.Price
}

func (t *WETH) SetContract(address string) {
	t.Contract = address
}
