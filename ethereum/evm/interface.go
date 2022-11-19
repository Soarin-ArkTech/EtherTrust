package evm

import (
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IExchangeLoader interface {
	IPubKeyGetter
	IPrivKeyGetter
	IEVMClientGetter
}

// Grabbers
type IPubKeyGetter interface {
	GetPubKey() *accounts.Account
}

type IPrivKeyGetter interface {
	GetPrivKey() *keystore.Key
}

type IEVMClientGetter interface {
	GetEVMClient() *ethclient.Client
}
