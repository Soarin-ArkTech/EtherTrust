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
func SendERC20(dreamTX ITokenTX) (string, error) {
	var tx EtherTXBuilder
	tx.SetRecipient("0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa") // wETH Token Contract
	tx.SetAmount(0)

	transferFnSignature := []byte("transfer(address,uint256)") // do not include spaces in the string

	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(dreamTX.GetWallet().Bytes(), 32)
	paddedAmount := common.LeftPadBytes(big.NewInt(int64(dreamTX.GetAmount())).Bytes(), 32)

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
