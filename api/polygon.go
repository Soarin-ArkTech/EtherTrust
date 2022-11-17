package etAPI

import "fmt"

func (api PolyscanTokenBal) GetBalWEI() {
	cbAPI := APICallBuilder{}
	cbAPI.SetMethod("GET")
	cbAPI.SetContentType("application/json")
	cbAPI.SetURL("https://api-testnet.polygonscan.com/api?module=account&action=tokenbalance&contractaddress=" + WETHContract + "&address=0xF5647Be44eA21d00240556A72672bEc75ed78D0A")

	coinbaseRes, err := cbAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to call out to Coinbase for ETH-USD spot in CallEthereum function. ", err)
	}

	_, err = ParseResults(coinbaseRes, &WrappedETH)
	if err != nil {
		fmt.Println("Failed to parse the CallEthereum function. ", err)
	}

}

func (api PolyscanTokenBal) SetToFloat32() float32 {
	return 1.0 // do l8r
}

type PolyscanTokenBal struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Result  *string `json:"result"`
}

var WrappedETH PolyscanTokenBal

const WETHContract = "0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa"
