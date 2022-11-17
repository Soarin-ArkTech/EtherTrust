package exchange

import "math/big"

type IBalanceReader interface {
	GetWEI() *big.Int
}
