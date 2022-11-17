package ether

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Our Keystore's Password (temporary, need to rewrite this file)
const secret string = "dapmeup1776"

func CreateKeys() *accounts.Account {
	// Create Directory
	err := os.Mkdir("keys", 0777)
	if err != nil {
		fmt.Printf("Unable to create Keys dir!\n%s", err)
	}

	// Set Keystore Working Directory & Encryption Algorithms
	ks := keystore.NewKeyStore("./keys", keystore.StandardScryptN, keystore.StandardScryptP)

	// Generate New Keypair
	account, err := ks.NewAccount(secret)
	if err != nil {
		log.Fatalln("We were unable to generate a new keypair for the exchange wallet! \n", err)
	}

	return &account
}

func (ether *Ethereum) LoadKeys() error {
	// Open Keystore or Output an Error if it is Missing
	keyFile, err := os.ReadFile("./keys/keys")
	if err != nil {
		CreateKeys()
	}

	// Set Keystore Working Directory & Encryption Algorithms
	ks := keystore.NewKeyStore("./keys/tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// Import Keypair to Memory as accounts.Account
	account, err := ks.Import(keyFile, secret, secret)
	if err != nil {
		fmt.Println("Unable to import the existing keys file!\n ", err)
	}

	// Clean Up After Func End
	err = os.RemoveAll("./keys/tmp")
	if err != nil {
		fmt.Println("Failed to remove temporary directory!")
	}

	// Grab our private keys to return too
	privKey, err := keystore.DecryptKey(keyFile, secret)
	if err != nil {
		fmt.Println("Failed to fetch private keys.")
	}

	// Assign to our Object
	EthereumClient.PubKey = &account
	EthereumClient.PrivKey = privKey

	return err
}
