package etAPI

import (
	"fmt"
	"math/big"
	"strconv"
)

func (api PolyscanTokenBal) QueryTokenBal() {
	cbAPI := APICallBuilder{}
	cbAPI.SetMethod("GET")
	cbAPI.SetContentType("application/json")
	cbAPI.SetURL(PolyscanAPI + "module=account&action=tokenbalance&contractaddress=" + WETHContract + "&address=0x16cde118c2acc7810591687156597f3bfb301193")

	coinbaseRes, err := cbAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to call out to Coinbase for ETH-USD spot in CallEthereum function. ", err)
	}

	_, err = ParseResults(coinbaseRes, &TreasuryWrappedETH)
	if err != nil {
		fmt.Println("Failed to parse the CallEthereum function. ", err)
	}

}

// Fetch WETH Wallet Bal
func (user PolyscanTokenBal) GetWEI() uint64 {
	wei, _ := strconv.Atoi(*user.Result)
	return big.NewInt(int64(wei)).Uint64()
}

// Turn
func BigToFloat32(bal *big.Float) float32 {
	bal32, _ := bal.Float32()
	return bal32
}

type PolyscanTokenBal struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Result  *string `json:"result"`
}

var TreasuryWrappedETH PolyscanTokenBal

const WETHContract = "0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa"
const PolyscanAPI = "https://api-testnet.polygonscan.com/api?"
