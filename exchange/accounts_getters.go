package exchange

import (
	"context"
	"fmt"
	"math"
	"math/big"

	"github.com/Soarin-ArkTech/ethereal-dreams/coinbase"

	ether "github.com/Soarin-ArkTech/ethereal-dreams/ethereum"
)

// Grab Player Balance in ETH
func (player ExchangeAccount) PlayerBalance() string {
	return fmt.Sprint(player.GetBalPow10())
}

// Grab Player Balane in USD
func (player ExchangeAccount) PlayerBalanceUSD() string {
	// bal, _ := ether.GetWalletBalance(player.Wallet).Float32()
	// ethprice, _ := strconv.ParseFloat(*coinbase.Ethereum.Amount, 32)

	bal, _ := player.GetBalPow10().Float32()

	return fmt.Sprint(bal * *coinbase.Ethereum.Amount)
}

// // Fetch EVM Wallet Bal
func (player ExchangeAccount) GetBalWEI() *big.Int {
	// Fetch raw balance
	weibal, err := ether.EthereumClient.Client.BalanceAt(context.Background(), player.Wallet, nil)
	if err != nil {
		fmt.Println("Failed to fetch balance of your wallet.")
	}

	return weibal
}

func (player ExchangeAccount) GetBalPow10() *big.Float {
	return WeiToNorm(player)
}

// Wei to 10^18 Decimal
func WeiToNorm(weiBal IBalanceReader) *big.Float {
	weiBigFloat, ok := new(big.Float).SetString(fmt.Sprint(weiBal.GetBalWEI()))
	if !ok {
		fmt.Println("Failed to make big float in WeiToNorm. ", ok)
	}

	return new(big.Float).Quo(weiBigFloat, big.NewFloat(math.Pow10(18)))
}

// To Raw Wei
func NormToWei(ether float32) uint64 {
	return uint64(ether * float32(math.Pow10(18)))
}
