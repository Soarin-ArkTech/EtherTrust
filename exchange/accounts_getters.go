package exchange

import (
	"context"
	"fmt"

	coinbaseAPI "github.com/Soarin-ArkTech/EtherTrust/api/coinbase"
	"github.com/Soarin-ArkTech/EtherTrust/ethereum/evm"
	"github.com/ethereum/go-ethereum/common"
)

// Grab Player Balance in ETH
func (player ExchangeAccount) GetPowAmount() float32 {
	ether, _ := evm.WeiToNorm(player).Float32()

	return ether
}

// Grab Player Balance in USD
func (player ExchangeAccount) GetUSD() float32 {
	return float32(player.GetPowAmount() * coinbaseAPI.Ethereum.CBToFloat32())
}

// Fetch EVM Wallet Bal
func (player ExchangeAccount) GetWEI() int64 {
	// Fetch raw WEI balance
	weibal, err := evm.EVMClient.Client.BalanceAt(context.Background(), player.Wallet, nil)
	if err != nil {
		fmt.Println("Failed to fetch balance of your wallet.")
	}

	return weibal.Int64()
}

func (player ExchangeAccount) GetWallet() common.Address {
	return player.Wallet
}

func (player ExchangeAccount) GetUUID() string {
	return player.UUID
}

func (player ExchangeAccountBuilder) GetUUID() string {
	return player.UUID
}
