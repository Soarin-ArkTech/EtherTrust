package evmkeys

import (
	"fmt"
	"os"

	"github.com/Soarin-ArkTech/EtherTrust/ethereum/evm"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

//////////////////////////////////////////////////////
///(temporary, need to rewrite this entire package)///
//////////////////////////////////////////////////////

// Our Keystore's Password (fake obv)
const secret string = "dapmeup1776"

func LoadKeys() error {
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
	evm.EVMClient.PubKey = &account
	evm.EVMClient.PrivKey = privKey

	return err
}
