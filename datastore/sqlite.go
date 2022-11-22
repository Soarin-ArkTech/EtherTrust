package etData

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type exchangeUser struct {
	UUID        string `gorm:"primary_key"`
	Wallet      common.Address
	FaucetPerms bool
}

type etSQLite struct {
	exchangeUser
	*gorm.DB
}

func (sqldb *etSQLite) LoadDB(file string) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		fmt.Println("Oh shit ", err)
	}

	sqldb.DB = db
}

func (sqldb etSQLite) GetUUID(uuid string) error {
	var result exchangeUser

	err := sqldb.DB.Find(&result, "UUID = ?", uuid).Error

	fmt.Println(result.UUID)

	return err
}
