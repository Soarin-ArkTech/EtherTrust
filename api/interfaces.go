package etAPI

import "math/big"

type IAPICaller interface {
}

type ITokenTicker interface {
	SetTicker() string
}

type IGetWEI interface {
	GetWEI() *big.Int
}
