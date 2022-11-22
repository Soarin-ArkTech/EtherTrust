package etData

import "github.com/ethereum/go-ethereum/common"

type DatabaseConn interface {
	GetUUID(uuid string) (string, error)
	CreateAccount(uuid string, wallet common.Address)
}

type ExchangeStore struct {
	db DatabaseConn
}
