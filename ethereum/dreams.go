package ether

import (
	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	"github.com/ethereum/go-ethereum/common"
)

type ITokenTX interface {
	GetPrice() float32
	GetAmount() int
	GetWallet() common.Address
}

type Dreams struct {
	Price float32
}

type DreamExchange struct {
	Dreams
	Amount int
	Wallet common.Address
}

func (d Dreams) GetPrice() float32 {
	return d.Price
}

func (d DreamExchange) GetAmount() int {
	return d.Amount
}
func (d DreamExchange) GetWallet() common.Address {
	return d.Wallet
}

func (d *Dreams) SetPrice(price float32) {
	d.Price = price / etAPI.Ethereum.CBToFloat32()
}

func (d *DreamExchange) SetWallet(wallet common.Address) {
	d.Wallet = wallet
}

func (d *DreamExchange) SetAmount(amnt int) {
	d.Amount = amnt
}
