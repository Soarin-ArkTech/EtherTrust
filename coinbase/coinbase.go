package coinbase

import (
	"fmt"

	"github.com/Soarin-ArkTech/ethereal-dreams/api"
)

func CallEthereum() error {
	cbAPI := api.APICallBuilder{}
	cbAPI.SetMethod("GET")
	cbAPI.SetURL("https://api.coinbase.com/v2/prices/ETH-USD/spot")
	cbAPI.SetContentType("application/json")
	coinbaseRes, err := cbAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to call out to Coinbase for ETH-USD spot in CallEthereum function. ", err)
	}

	defer coinbaseRes.Body.Close()

	_, err = api.ParseResults(coinbaseRes, &Ethereum)
	if err != nil {
		fmt.Println("Failed to parse the CallEthereum function. ", err)
	}

	return err
}

func (Coinbase) CoinbaseToFloat32() float32 {
	// Convert string into foat64
	// eth, err := strconv.ParseFloat(*Ethereum.Amount, 32)
	// if err != nil {
	// 	fmt.Println("Failed to convert Ethereum output from string to float!")
	// }

	return float32(*Ethereum.Amount)
}

var Ethereum Coinbase

type Coinbase struct {
	Crypto `json:"data"`
}

type Crypto struct {
	BaseToken     string   `json:"base"`
	ComparedToken string   `json:"currency"`
	Amount        *float32 `json:"amount"`
}
