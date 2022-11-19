package evmkeys

import (
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

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
