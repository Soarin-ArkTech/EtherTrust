package evm

import (
	"fmt"
	"math"
	"math/big"

	ether "github.com/Soarin-ArkTech/EtherTrust/ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Wei to 10^18 Decimal
func WeiToNorm(weiBal ether.IWEIGetter) *big.Float {
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

func (e EVM) GetPrivKey() *keystore.Key {
	return e.PrivKey
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
