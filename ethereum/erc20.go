package ether

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type Token struct {
	Price    float32
	Contract string
}

type TokenExchange struct {
	Token
	Amount float32
	Wallet common.Address
}

// Sends ERC20 to Destination
func TransferERC20(req ITokenTX) (string, bool) {
	var contractData []byte
	var tx EtherTXBuilder
	tx.SetRecipient(req.GetContract()) // wETH Token Contract
	tx.SetAmount(0)

	// Our Data to Send to Smart Contract
	methodID := GetMethodID("transfer(address,uint256)")
	paddedAddress := PadAddress(req)
	paddedAmount := PadAmount(req)
	contractData = append(contractData, methodID...)
	contractData = append(contractData, paddedAddress...)
	contractData = append(contractData, paddedAmount...)
	tx.SetData(contractData)

	// Build & Send TX to Blockchain
	sentTX, ok := BroadcastTX(tx.BuildTX())

	return fmt.Sprintf("tx sent: %s\n", sentTX.Hash().Hex()), ok
}

func GetMethodID(method string) []byte {
	funcSignature := []byte(method) // do not include spaces in the string
	hash := sha3.NewLegacyKeccak256()
	hash.Write(funcSignature)

	return hash.Sum(nil)[:4]
}

func PadAddress(token ITokenTX) []byte {
	return common.LeftPadBytes(token.GetWallet().Bytes(), 32)
}

func PadAmount(token ITokenTX) []byte {
	return common.LeftPadBytes(big.NewInt(int64(token.GetWEI())).Bytes(), 32)
}
