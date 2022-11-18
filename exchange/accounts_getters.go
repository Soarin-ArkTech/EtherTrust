package exchange

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	etAPI "github.com/Soarin-ArkTech/ethereal-dreams/api"
	ether "github.com/Soarin-ArkTech/ethereal-dreams/ethereum"
)

// Grab Player Balance in ETH
func (player ExchangeAccount) WalletBalance() string {
	return fmt.Sprint(player.GetBalPow10())
}

// Grab Player Balane in USD
func (player ExchangeAccount) WalletBalanceUSD() float32 {
	// bal, _ := ether.GetWalletBalance(player.Wallet).Float32()
	ethprice, _ := strconv.ParseFloat(*etAPI.Ethereum.Amount, 32)

	bal, _ := player.GetBalPow10().Float64()

	return float32(bal * ethprice)
}

// // Fetch EVM Wallet Bal
func (player ExchangeAccount) GetWEI() *big.Int {
	// Fetch raw balance
	weibal, err := ether.EthereumClient.Client.BalanceAt(context.Background(), player.Wallet, nil)
	if err != nil {
		fmt.Println("Failed to fetch balance of your wallet.")
	}

	return weibal
}

func (player ExchangeAccount) GetBalPow10() *big.Float {
	return ether.WeiToNorm(player)
}
