package exchange

import "math/big"

type IBalanceReader interface {
	GetBalWEI() *big.Int
}
