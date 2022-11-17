package coinbase

import (
	"fmt"
	"strconv"

	apiclient "github.com/Soarin-ArkTech/ethereal-dreams/api"
)

func CallEthereum() error {
	cbAPI := apiclient.APICallBuilder{}
	cbAPI.SetMethod("GET")
	cbAPI.SetContentType("application/json")
	cbAPI.SetURL("https://api.coinbase.com/v2/prices/ETH-USD/spot")

	coinbaseRes, err := cbAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to call out to Coinbase for ETH-USD spot in CallEthereum function. ", err)
	}

	_, err = apiclient.ParseResults(coinbaseRes, &Ethereum)
	if err != nil {
		fmt.Println("Failed to parse the CallEthereum function. ", err)
	}

	return err
}

func (Coinbase) CoinbaseToFloat32() float32 {
	// Convert string into foat64
	eth, err := strconv.ParseFloat(*Ethereum.Amount, 32)
	if err != nil {
		fmt.Println("Failed to convert Ethereum output from string to float!")
	}

	return float32(eth)
}

var Ethereum Coinbase

type Coinbase struct {
	Crypto `json:"data"`
}

type Crypto struct {
	BaseToken     string  `json:"base"`
	ComparedToken string  `json:"currency"`
	Amount        *string `json:"amount"`
}
