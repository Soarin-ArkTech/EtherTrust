package ether

import (
	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	"github.com/ethereum/go-ethereum/common"
)

type ITokenTXReader interface {
	GetWEI() uint64
	GetWallet() common.Address
	GetContract() string
}

type Dreams struct {
	Price float32
}

type DreamExchange struct {
	Dreams
	Amount   float32
	Wallet   common.Address
	Contract string
}

func (d DreamExchange) GetPrice() float32 {
	return d.Price
}

func (d DreamExchange) GetAmount() float32 {
	return d.Amount
}

func (d DreamExchange) GetWEI() uint64 {
	return NormToWei(d.Amount)
}

func (d DreamExchange) GetWallet() common.Address {
	return d.Wallet
}

func (d DreamExchange) GetContract() string {
	return d.Contract
}

func (d DreamExchange) GetUSD() float32 {
	return d.Amount * etAPI.Ethereum.CBToFloat32()
}

func (d *DreamExchange) SetPrice(price float32) {
	d.Price = price / etAPI.Ethereum.CBToFloat32()
}

func (d *DreamExchange) SetWallet(wallet common.Address) {
	d.Wallet = wallet
}

func (d *DreamExchange) SetAmount(amnt int) {
	d.Amount = float32(amnt) * d.Price
}

func (d *DreamExchange) SetContract(address string) {
	d.Contract = "0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa"
}
