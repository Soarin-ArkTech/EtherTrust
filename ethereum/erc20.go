package ether

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
)

// Sends ETH to Destination
// func SendUNI(wallet common.Address, amount uint64) (string, error) {
func TransferERC20(req ITokenTXReader) (string, error) {
	var tx EtherTXBuilder
	tx.SetRecipient(req.GetContract()) // wETH Token Contract
	tx.SetAmount(0)

	var methodID []byte
	var paddedAddress []byte
	var paddedAmount []byte

	methodID = GetMethodID("transfer(address,uint256)")
	paddedAddress = PadAddress(req)
	paddedAmount = PadAmount(req)

	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	tx.SetData(data)

	fmt.Println("Address is ", hexutil.Encode(paddedAddress))
	fmt.Println("Amount is ", hexutil.Encode(paddedAmount))

	fmt.Printf("\nTo %q, Gas %v, GasLimit %v\n", tx.RecipientWallet, tx.GasPrice, tx.GasLimit)

	// fmt.Println(tx.RecipientWallet, tx.Amount, tx.Data)

	sentTX, err := BroadcastTX(tx.BuildTX())
	if err != nil {
		fmt.Println("Failed to post the Ethereum transaction. Error: ", err)
	}

	return fmt.Sprintf("tx sent: %s\n", sentTX.Hash().Hex()), err
}

func GetMethodID(method string) []byte {
	funcSignature := []byte(method) // do not include spaces in the string
	hash := sha3.NewLegacyKeccak256()
	hash.Write(funcSignature)

	return hash.Sum(nil)[:4]
}

func PadAddress(token ITokenTXReader) []byte {
	return common.LeftPadBytes(token.GetWallet().Bytes(), 32)
}

func PadAmount(token ITokenTXReader) []byte {
	return common.LeftPadBytes(big.NewInt(int64(token.GetWEI())).Bytes(), 32)
}
