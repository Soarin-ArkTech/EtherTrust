package evm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IExchangeLoader interface {
	IPubKeyAddressGetter
	IPrivKeyAddressGetter
	IEVMClientGetter
}

// Composed Interfaces
type EVMGetter interface {
	ICoinTX
	INonceGetter
	IGasGetter
	IChainGetter
}

type IEVMTX interface {
	IRawTXGetter
	EVMGetter
	IPrivKeyAddressGetter
}

type ICoinTX interface {
	IWEIGetter
	IWalletGetter
}

// Segmented Interfaces
type IWEIGetter interface {
	GetWEI() uint64
}

type IWalletGetter interface {
	GetWallet() common.Address
}

type INonceGetter interface {
	GetNonce() uint64
}

type IGasGetter interface {
	GetGasPrice() *big.Int
	GetGasLimit() uint64
}

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

type IChainGetter interface {
	GetChainID() *big.Int
}

type IRawTXGetter interface {
	GetRawTX() *types.Transaction
}
