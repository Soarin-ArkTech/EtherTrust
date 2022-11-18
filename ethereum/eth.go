package ether

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Create Connection to our RPC
func (ether *Ethereum) DialRPC() {
	var err error
	EthereumClient.Client, err = ethclient.Dial("https://rpc.ankr.com/polygon_mumbai")
	if err != nil {
		fmt.Println("Could not connect to RPC! Err: ", err)
	}

	//"https://rpc.ankr.com/eth_goerli"
}

// Sends ETH to Destination
func SendETH(wallet string, amount uint64) (string, error) {
	var tx EtherTXBuilder
	tx.SetRecipient(wallet)
	tx.SetAmount(amount)

	sentTX, err := BroadcastTX(tx.BuildTX())
	if err != nil {
		fmt.Println("Failed to post the Ethereum transaction. Error: ", err)
	}

	return fmt.Sprintf("tx sent: %s\n", sentTX.Hash().Hex()), err
}

// Broadcast to Blockchain
func BroadcastTX(ethertx IUnsignedTX) (*types.Transaction, error) {
	tx := SignTX(ethertx)

	err := EthereumClient.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Printf("Failed to broadcast TX. Error: %q\n", err)
	} else {
		IncrementNonce()
	}

	return tx, err
}

// Sign an Unsigned Transaction
func SignTX(unsignedTX IUnsignedTX) *types.Transaction {
	tx, err := types.SignTx(unsignedTX.GetRawTX(), types.NewEIP155Signer(unsignedTX.GetChainID()), EthereumClient.PrivKey.PrivateKey)
	if err != nil {
		fmt.Printf("Failed to sign the transaction with the exchange private key. Error: %q\n", err)
	}

	return tx
}

// Get Account's Current Nonce
func (ether *Ethereum) GetAccountNonce() uint64 {
	nonce, err := ether.Client.NonceAt(context.Background(), EthereumClient.PubKey.Address, nil)
	if err != nil {
		fmt.Println("Unable to send transaction due to Nonce error! ", err)
	}

	fmt.Println("GetAccountNonce result is ", nonce)
	return nonce
}

// Get Account's Next Nonce and validate the output
func (ether *Ethereum) GetPendingNonce() uint64 {
	CheckNonce := func() *uint64 {
		pNonce, _ := EthereumClient.Client.PendingNonceAt(context.Background(), EthereumClient.PubKey.Address)
		return &pNonce
	}

	for *CheckNonce() > 50000 {
		fmt.Println("Pending Nonce fucked!")
		CheckNonce()
		time.Sleep(time.Second * 3)
	}

	return *CheckNonce()
}

// Wei to 10^18 Decimal
func WeiToNorm(weiBal IBalanceReader) *big.Float {
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

var EthereumClient Ethereum

type Ethereum struct {
	PubKey  *accounts.Account
	PrivKey *keystore.Key
	Client  *ethclient.Client
}
