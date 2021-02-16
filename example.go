package bittrex

import (
	"log"
	"net/http"
	"os"

	"github.com/shopspring/decimal"
)

func example() {
	uri := "https://api.bittrex.com/v3"
	apiKey := os.Getenv("BITTREX_API_KEY")
	secretKey := os.Getenv("BITTREX_SECRET_KEY")
	httpClient := &http.Client{}
	client := NewBittrexClient(apiKey, secretKey, httpClient)
	api := NewBittrexAPI(client, uri)
	//log.Println(api.getCurrency("BTC"))
	//log.Println(api.getBalances())
	//log.Printf("getMarket: %v", api.getMarket("ETH-BTC"))
	//log.Printf("getMarketSummary: %v", api.getMarketSummary("ETH-BTC"))
	//log.Printf("getMarketSummaries: %v", api.getMarketSummaries())
	//log.Printf("getMarketTicker: %v", api.getMarketTicker("ETH-BTC"))
	//log.Printf("getORders: %v,", api.getOrders("Open"))
	//orders, _ := api.getOrders("Closed")
	//Required marketSymbol, direction, type, timeInForce

	decimalValue := decimal.NewFromFloat(1)
	order := Order{
		OrderID:      "",
		MarketSymbol: "ETH-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		TimeInForce:  "GOOD_TIL_CANCELLED",
		Quantity:     &decimalValue,
		Limit:        &decimalValue,
	}

	response, err := api.CreateOrder(order)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("response: %v", response)
}
