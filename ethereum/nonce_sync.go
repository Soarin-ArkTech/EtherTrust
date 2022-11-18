package ether

import (
	"fmt"
)

var SeqNonce *uint64

// Increment Local Counter
func IncrementNonce() {
	*SeqNonce++
}

// Sync Local Nonce Counter with Ethereum Nonce Counter
func NonceSync() string {
	nonce := EthereumClient.GetAccountNonce()

	if EthereumClient.GetPendingNonce()-nonce >= 10 {
		SeqNonce = &nonce
	} else {
		pendingNonce := EthereumClient.GetPendingNonce()
		SeqNonce = &pendingNonce
	}

	return fmt.Sprintf("delta: %v, ", EthereumClient.GetPendingNonce()-nonce)
}
