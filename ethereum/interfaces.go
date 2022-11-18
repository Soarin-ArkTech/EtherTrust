package ether

import "math/big"

type IBalanceReader interface {
	GetWEI() *big.Int
}
