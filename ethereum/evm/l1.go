package evm

import (
	"context"
	"fmt"
	"time"

	ether "github.com/Soarin-ArkTech/EtherTrust/ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

// Sends MATIC to Destination
func TransferCoin(req ether.ICoinTX) (string, bool) {
	var tx EtherTXBuilder
	tx.SetWallet(req.GetWallet().Hex())
	tx.SetAmount(req.GetWEI())

	// Build & Send TX to Blockchain
	sentTX, ok := BroadcastTX(tx.BuildTX())

	return fmt.Sprintf("tx sent: %s\n", sentTX.Hash().Hex()), ok
}

// Broadcast to Blockchain
func BroadcastTX(ethertx ether.IEVMTX) (*types.Transaction, bool) {
	tx := SignTX(ethertx)

	err := EVMClient.Client.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Printf("Failed to broadcast TX. Error: %q\n", err)
		DecrementNonce()
		return nil, false
	}

	return tx, true
}

// Sign an Unsigned Transaction
func SignTX(unsignedTX ether.IEVMTX) *types.Transaction {
	tx, err := types.SignTx(unsignedTX.GetRawTX(), types.NewEIP155Signer(unsignedTX.GetChainID()), EVMClient.PrivKey.PrivateKey)
	if err != nil {
		fmt.Printf("Failed to sign the transaction with the exchange private key. Error: %q\n", err)
	}

	return tx
}

// Get Account's Current Nonce
func (evm *EVM) GetAccountNonce() uint64 {
	nonce, err := evm.Client.NonceAt(context.Background(), evm.GetPubKeyAddress() nil)
	if err != nil {
		fmt.Println("Unable to send transaction due to Nonce error! ", err)
	}

	fmt.Println("GetAccountNonce result is ", nonce)
	return nonce
}

// Get Account's Next Nonce and validate the output
func (evm *EVM) GetPendingNonce() uint64 {
	CheckNonce := func() *uint64 {
		pNonce, _ := EVMClient.Client.PendingNonceAt(context.Background(), evm.GetPubKeyAddress())
		return &pNonce
	}

	for *CheckNonce() > 50000 {
		fmt.Println("Pending Nonce fucked!")
		CheckNonce()
		time.Sleep(time.Second * 3)
	}

	return *CheckNonce()
}
