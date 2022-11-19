package exchange

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

const ExchangeStore string = "plswork.json"

// Load Exchange Users and Return Slice of Exchange Accounts to Add Onto
func LoadAccounts() ([]ExchangeAccountBuilder, error) {
	Accounts := []ExchangeAccountBuilder{}

	accBytes, err := os.ReadFile("plswork.json")
	if err != nil {
		fmt.Printf("We're unable to read the exchange wallet datastore.\n Error: %s", err)
	}

	err = json.Unmarshal(accBytes, &Accounts)
	if err != nil {
		fmt.Printf("Failed to load wallets into memory.\n Error: %s", err)
	}

	return Accounts, err
}

// Select User from Exchange Store
func LoadUser(player proxy.Player) ExchangeAccount {
	accounts, _ := LoadAccounts()
	var user *ExchangeAccount

	for _, playerIDs := range accounts {
		if playerIDs.UUID == player.ID().String() {
			user = &ExchangeAccount{
				UUID:   playerIDs.UUID,
				Wallet: playerIDs.Wallet,
			}
		}
	}

	return *user
}

// Save User to Exchange store
func WriteAccount(input []ExchangeAccountBuilder) error {

	output, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		fmt.Println("Failed to encode the accounts in memory to JSON!")
	}

	err = os.WriteFile("plswork.json", output, 0644)
	if err != nil {
		fmt.Println("Failed to write to the exchange filestore json.")
	}

	return err
}

// Overwrite Wallet if User Exists
func OverwriteAccount(db []ExchangeAccountBuilder, dupeIndex int) []ExchangeAccountBuilder {
	return append(db[:dupeIndex], db[dupeIndex+1:]...)
}

// Check if User Exists or Not
func CheckAccountsDB(db []ExchangeAccountBuilder, userMap map[string]string) (int, bool) {
	for k, v := range db {
		_, exists := userMap[v.UUID]
		if exists {
			fmt.Printf("User already exists, modifying index %v.", k)
			return k, true
		}
	}

	return 0, false
}

// Convert Structs into Map to Check if User Exists
func (input ExchangeAccountBuilder) StructToMap() map[string]string {
	account := make(map[string]string)
	account[input.UUID] = input.Wallet.String()

	return account
}

// Build User and Add To Slice of Users
func (user ExchangeAccountBuilder) Build(db []ExchangeAccountBuilder) {
	userMap := user.StructToMap()
	duped, existing := CheckAccountsDB(db, userMap)

	if existing {
		modifiedUser := OverwriteAccount(db, duped)
		WriteAccount(append(modifiedUser, user))
	} else {
		addUser := append(db, user)
		WriteAccount(addUser)
		GiveFaucetPerm(user.UUID)
	}
}

type ExchangeAccount struct {
	UUID   string         `json:"uuid"`
	Wallet common.Address `json:"wallet"`
}

type ExchangeAccountBuilder struct {
	UUID   string         `json:"uuid"`
	Wallet common.Address `json:"wallet"`
}
