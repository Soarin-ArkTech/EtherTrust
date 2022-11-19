package erc20

import (
	"fmt"
	"math/big"

	"github.com/Soarin-ArkTech/EtherTrust/ethereum/evm"
	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

type Token struct {
	Price    float32 // Exchange Rate for 1 Ingame Dream
	Contract string  // Specify Token Contract
}

type TokenExchange struct {
	Token
	PowAmount float32 // Amount in 10^18 decimal
	Wallet    common.Address
}

// Sends ERC20 to Destination
func TransferERC20(req ITokenTX) (string, bool) {
	var tx evm.EtherTXBuilder
	tx.SetWallet(req.GetContract()) // wETH Token Contract
	tx.SetAmount(0)

	// Our Data to Send to Smart Contract
	methodID := GetMethodID("transfer(address,uint256)")
	tx.SetData(ContractCaller(methodID, req))

	// Build & Send TX to Blockchain
	sentTX, ok := evm.BroadcastTX(tx.BuildTX())

	return fmt.Sprintf("tx sent: %s\n", sentTX.Hash().Hex()), ok
}

func GetMethodID(method string) []byte {
	funcSignature := []byte(method) // do not include spaces in the string
	hash := sha3.NewLegacyKeccak256()
	hash.Write(funcSignature)

	return hash.Sum(nil)[:4]
}

// Builds & Pads our Data to Attach to Transaction
func ContractCaller(methodSignature []byte, req ITokenTX) []byte {
	var contractData []byte
	recipient := PadAddress(req)
	amount := PadAmount(req)

	contractData = append(contractData, methodSignature...)
	contractData = append(contractData, recipient...)
	contractData = append(contractData, amount...)

	return contractData
}

func PadAddress(token evm.IWalletGetter) []byte {
	return common.LeftPadBytes(token.GetWallet().Bytes(), 32)
}

func PadAmount(token evm.IWEIGetter) []byte {
	return common.LeftPadBytes(big.NewInt(int64(token.GetWEI())).Bytes(), 32)
}
