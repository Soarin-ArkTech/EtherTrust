package ether

import (
	"fmt"

	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	"github.com/ethereum/go-ethereum/common"
)

type ITokenTX interface {
	GetPrice() float32
	GetAmount() uint64
	GetWallet() common.Address
}

type Dreams struct {
	Price float32
}

type DreamExchange struct {
	Dreams
	Amount uint64
	Wallet common.Address
}

func (d DreamExchange) GetPrice() float32 {
	return d.Price
}

func (d DreamExchange) GetAmount() uint64 {
	return d.Amount
}

func (d DreamExchange) GetWallet() common.Address {
	return d.Wallet
}

func (d DreamExchange) GetUSD() float32 {
	return float32(d.Amount) * etAPI.Ethereum.CBToFloat32()
}

func (d *DreamExchange) SetPrice(price float32) {
	d.Price = price / etAPI.Ethereum.CBToFloat32()
}

func (d *DreamExchange) SetWallet(wallet common.Address) {
	d.Wallet = wallet
}

func (d *DreamExchange) SetAmount(amnt int) {

	fmt.Printf("Amount: %v\nPrice: %f\n", amnt, d.Price)
	d.Amount = NormToWei(float32(amnt) * d.Price)
}
