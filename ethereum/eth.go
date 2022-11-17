package ether

import (
	"context"
	"fmt"
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

// // Fetch EVM Wallet Bal
// func GetWalletBalance(wallet common.Address) *big.Float {
// 	// Fetch raw balance
// 	weibal, err := EthereumClient.Client.BalanceAt(context.Background(), wallet, nil)
// 	if err != nil {
// 		fmt.Println("Failed to fetch balance of your wallet.")
// 	}

// 	return WeiToNorm(weibal)
// }

// func GetTokenBalance(wallet string, contract string) *big.Int {
// 	// Fetch raw balance
// 	weibal, err := EthereumClient.Client.CallContract(context.Background(), ethereum.CallMsg{To: (*common.Address)(common.HexToAddress(wallet).Bytes())})
// 	if err != nil {
// 		fmt.Println("Failed to fetch balance of your wallet.")
// 	}

// 	return WeiToNorm(weibal)
// }

var EthereumClient Ethereum

type Ethereum struct {
	PubKey  *accounts.Account
	PrivKey *keystore.Key
	Client  *ethclient.Client
}
