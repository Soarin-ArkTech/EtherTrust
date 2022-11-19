package evm

import (
	"fmt"
)

var SeqNonce *uint64

// Increment Local Counter
func IncrementNonce() {
	*SeqNonce++
}

// Decrement Local Counter
func DecrementNonce() {
	*SeqNonce--
}

// Sync Local Nonce Counter with Ethereum Nonce Counter
func NonceSync() string {
	nonce := EVMClient.GetAccountNonce()

	if EVMClient.GetPendingNonce()-nonce >= 10 {
		SeqNonce = &nonce
	} else {
		pendingNonce := EVMClient.GetPendingNonce()
		SeqNonce = &pendingNonce
	}

	return fmt.Sprintf("delta: %v, ", EVMClient.GetPendingNonce()-nonce)
}
