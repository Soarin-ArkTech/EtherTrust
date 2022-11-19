package evm

import (
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// // Wei to 10^18 Decimal
func WeiToNorm(weiBal IWEIGetter) *big.Float {
	weiBigFloat, ok := new(big.Float).SetString(fmt.Sprint(weiBal.GetWEI()))
	if !ok {
		fmt.Println("Failed to make big float in WeiToNorm. ", ok)
	}

	return new(big.Float).Quo(weiBigFloat, big.NewFloat(math.Pow10(18)))
}

// To Raw Wei
func NormToWei(ether float32) uint64 {
	return uint64(ether * float32(math.Pow10(18)))
}

func (e EVM) GetPubKey() *accounts.Account {
	return e.PubKey
}

func (e EVM) GetPubKeyAddress() common.Address {
	return e.PubKey.Address
}

func (e EVM) GetPrivKey() *ecdsa.PrivateKey {
	return e.PrivKey.PrivateKey
}

func (e EVM) GetPrivKeyAddress() common.Address {
	return e.PubKey.Address
}

func (e EVM) GetEVMClient() *ethclient.Client {
	return e.Client
}

var EVMClient EVM

type EVM struct {
	PubKey  *accounts.Account
	PrivKey *keystore.Key
	Client  *ethclient.Client
}
