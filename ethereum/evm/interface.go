package evm

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IExchangeLoader interface {
	IPubKeyAddressGetter
	IPrivKeyAddressGetter
	IEVMClientGetter
}

// Grabbers
type IPubKeyGetter interface {
	GetPubKey() *accounts.Account
}

type IPubKeyAddressGetter interface {
	GetPubKeyAddress() common.Address
}

type IPrivKeyGetter interface {
	GetPrivKey() *keystore.Key
}

type IPrivKeyAddressGetter interface {
	GetPrivKeyAddress() common.Address
}

type IEVMClientGetter interface {
	GetEVMClient() *ethclient.Client
}
