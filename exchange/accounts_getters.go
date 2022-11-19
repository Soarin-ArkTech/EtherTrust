package exchange

import (
	"context"
	"fmt"
	"strconv"

	etAPI "github.com/Soarin-ArkTech/EtherTrust/api"
	ether "github.com/Soarin-ArkTech/EtherTrust/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Grab Player Balance in ETH
func (player ExchangeAccount) GetPowAmount() float32 {
	ether, _ := ether.WeiToNorm(player).Float32()

	return ether
}

// Grab Player Balance in USD
func (player ExchangeAccount) GetUSD() float32 {
	ethprice, _ := strconv.ParseFloat(*etAPI.Ethereum.Amount, 32)

	return float32(player.GetPowAmount() * float32(ethprice))
}

// Fetch EVM Wallet Bal
func (player ExchangeAccount) GetWEI() uint64 {
	// Fetch raw WEI balance
	weibal, err := ether.EthereumClient.Client.BalanceAt(context.Background(), player.Wallet, nil)
	if err != nil {
		fmt.Println("Failed to fetch balance of your wallet.")
	}

	return weibal.Uint64()
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
