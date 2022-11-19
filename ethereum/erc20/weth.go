package erc20

import (
	coinbaseAPI "github.com/Soarin-ArkTech/EtherTrust/api/coinbase"
	"github.com/Soarin-ArkTech/EtherTrust/ethereum/evm"
	"github.com/ethereum/go-ethereum/common"
)

type WETH TokenExchange

func (t WETH) GetSpot() float32 {
	return t.Price
}

func (t WETH) GetPowAmount() float32 {
	return t.PowAmount
}

func (t WETH) GetWEI() int64 {
	return evm.NormToWei(t.PowAmount) // WEI Amount to Send
}

func (t WETH) GetWallet() common.Address {
	return t.Wallet
}

func (t WETH) GetContract() string {
	return t.Contract
}

func (t WETH) GetUSD() float32 {
	return t.PowAmount * coinbaseAPI.Ethereum.CBToFloat32()
}

func (t *WETH) SetPrice(price float32) {
	t.Price = price / coinbaseAPI.Ethereum.CBToFloat32()
}

func (t *WETH) SetWallet(wallet common.Address) {
	t.Wallet = wallet
}

func (t *WETH) SetAmount(amnt int) {
	t.PowAmount = float32(amnt) * t.Price
}

func (t *WETH) SetContract(address string) {
	t.Contract = address
}
