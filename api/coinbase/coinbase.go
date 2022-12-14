package coinbaseAPI

import (
	"fmt"
	"strconv"

	etAPI "github.com/Soarin-ArkTech/EtherTrust/api/client"
)

func (CBSpot) GetSpot() {
	cbAPI := etAPI.APICallBuilder{}
	cbAPI.SetMethod("GET")
	cbAPI.SetContentType("application/json")
	cbAPI.SetURL("https://api.coinbase.com/v2/prices/ETH-USD/spot")

	coinbaseRes, err := cbAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to call out to Coinbase for ETH-USD spot in CallEthereum function. ", err)
	}

	etAPI.ParseResults(coinbaseRes, &Ethereum)
}

func (CBSpot) CBToFloat32() float32 {
	// Convert string into float64
	eth, err := strconv.ParseFloat(*Ethereum.Amount, 32)
	if err != nil {
		fmt.Println("Failed to convert Ethereum output from string to float!")
	}

	return float32(eth)
}

var Ethereum CBSpot

type CBSpot struct {
	CBPriceSpot `json:"data"`
}

type CBPriceSpot struct {
	BaseToken     string  `json:"base"`
	ComparedToken string  `json:"currency"`
	Amount        *string `json:"amount"`
}
