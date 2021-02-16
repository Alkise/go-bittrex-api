package bittrex

import (
	"errors"
	"net/http"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestBittrexAPIFixture(t *testing.T) {
	gunit.Run(new(BittrexAPIFixture), t)

}

type BittrexAPIFixture struct {
	*gunit.Fixture
}

func (this *BittrexAPIFixture) Setup() {}

func (this *BittrexAPIFixture) TestGetCurrency() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetCurrency("fakesymbol")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, Currency{})
}

func (this *BittrexAPIFixture) TestGetBalances() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetBalances()
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []Balance{
		{
			CurrencySymbol: "BTC",
			Total:          "0.00000000",
			Available:      "0.00000000",
			UpdatedAt:      "2019-10-29T20:25:10.16Z",
		}, {
			CurrencySymbol: "LTC",
			Total:          "0",
			Available:      "0",
			UpdatedAt:      "2020-09-03T21:27:53.8210894Z",
		},
	})
}
func (this *BittrexAPIFixture) TestGetMarket() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarket("fakesymbol")
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, Market{
		Symbol:              "ETH-BTC",
		BaseCurrencySymbol:  "ETH",
		QuoteCurrencySymbol: "BTC",
		MinTradeSize:        "0.01000000",
		Precision:           8,
		Status:              "ONLINE",
		CreatedAt:           "2015-08-14T09:02:24.817Z",
		Notice:              "",
		ProhibitedIn:        []string{},
	})
}

func (this *BittrexAPIFixture) TestGetMarkets() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarkets()
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []Market{
		{
			Symbol:              "4ART-BTC",
			BaseCurrencySymbol:  "4ART",
			QuoteCurrencySymbol: "BTC",
			MinTradeSize:        "10.00000000",
			Precision:           8,
			Status:              "ONLINE",
			CreatedAt:           "2020-06-10T15:05:29.833Z",
			Notice:              "",
			ProhibitedIn:        []string{"US"},
		}, {
			Symbol:              "4ART-USDT",
			BaseCurrencySymbol:  "4ART",
			QuoteCurrencySymbol: "USDT",
			MinTradeSize:        "10.00000000",
			Precision:           5,
			Status:              "ONLINE",
			CreatedAt:           "2020-06-10T15:05:40.98Z",
			Notice:              "",
			ProhibitedIn:        []string{"US"},
		},
	})
}

func (this *BittrexAPIFixture) TestGetMarketSummary() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarketSummary("fakesymbol")
	high := decimal.NewFromFloatWithExponent(0.03894964, -8)
	low := decimal.NewFromFloatWithExponent(0.03650000, -8)
	volume := decimal.NewFromFloat(18494.04035144)
	quoteVolume := decimal.NewFromFloat(696.42899671)
	percentChange := decimal.NewFromFloat(-3.33)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, MarketSummary{
		Symbol:        "ETH-BTC",
		High:          &high,
		Low:           &low,
		Volume:        &volume,
		QuoteVolume:   &quoteVolume,
		PercentChange: &percentChange,
		UpdatedAt:     "2020-09-04T04:37:45.107Z",
	})
}

func (this *BittrexAPIFixture) TestGetMarketSummaries() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarketSummaries()
	firstInstrumentHigh := decimal.NewFromFloat(0.00000275)
	firstInstrumentLow := decimal.NewFromFloat(0.00000249)
	firstInstrumentVolume := decimal.NewFromFloat(54499.59344453)
	firstInstrumentQuoteVolume := decimal.NewFromFloat(0.13917073)
	firstInstrumentPercentChange := decimal.NewFromFloat(10.44)
	secondInstrumentHigh := decimal.NewFromFloatWithExponent(0.02880000, -8)
	secondInstrumentLow := decimal.NewFromFloatWithExponent(0.02667000, -8)
	secondInstrumentVolume := decimal.NewFromFloat(48259.53706735)
	secondInstrumentQuoteVolume := decimal.NewFromFloat(1320.75839607)
	secondInstrumentPercentChange := decimal.NewFromFloat(-6.11)
	this.So(err, should.BeNil)

	this.So(result, should.Resemble, []MarketSummary{
		{
			Symbol:        "4ART-BTC",
			High:          &firstInstrumentHigh,
			Low:           &firstInstrumentLow,
			Volume:        &firstInstrumentVolume,
			QuoteVolume:   &firstInstrumentQuoteVolume,
			PercentChange: &firstInstrumentPercentChange,
			UpdatedAt:     "2020-09-04T04:58:55.447Z",
		}, {
			Symbol:        "4ART-USDT",
			High:          &secondInstrumentHigh,
			Low:           &secondInstrumentLow,
			Volume:        &secondInstrumentVolume,
			QuoteVolume:   &secondInstrumentQuoteVolume,
			PercentChange: &secondInstrumentPercentChange,
			UpdatedAt:     "2020-09-04T04:33:20.01Z",
		},
	})
}

func (this *BittrexAPIFixture) TestGetMarketTicker() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarketTicker("fakesymbol")
	lastTradeRate := decimal.NewFromFloat(0.03760069)
	bidRate := decimal.NewFromFloat(0.03760103)
	askRate := decimal.NewFromFloat(0.03762798)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, MarketTicker{
		Symbol:        "ETH-BTC",
		LastTradeRate: &lastTradeRate,
		BidRate:       &bidRate,
		AskRate:       &askRate,
	})
}

func (this *BittrexAPIFixture) TestGetMarketTickers() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetMarketTickers()
	firstInstrumentLastTradeRate := decimal.NewFromFloat(0.03760069)
	firstInstrumentBidRate := decimal.NewFromFloat(0.03760103)
	firstInstrumentAskRate := decimal.NewFromFloat(0.03762798)
	secondInstrumentLastTradeRate := decimal.NewFromFloat(1.03760069)
	secondInstrumentBidRate := decimal.NewFromFloat(1.03760103)
	secondInstrumentAskRate := decimal.NewFromFloat(1.03762798)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []MarketTicker{
		{
			Symbol:        "ETH-BTC",
			LastTradeRate: &firstInstrumentLastTradeRate,
			BidRate:       &firstInstrumentBidRate,
			AskRate:       &firstInstrumentAskRate,
		},
		{
			Symbol:        "ETH-FAKE",
			LastTradeRate: &secondInstrumentLastTradeRate,
			BidRate:       &secondInstrumentBidRate,
			AskRate:       &secondInstrumentAskRate,
		},
	})
}

func (this *BittrexAPIFixture) TestGetOrder() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetOrder("fakeOrder")
	quantity := decimal.NewFromFloatWithExponent(77.53046131, -8)
	limit := decimal.NewFromFloatWithExponent(0.00003528, -8)
	fillQuantity := decimal.NewFromFloatWithExponent(77.53046131, -8)
	commission := decimal.NewFromFloatWithExponent(0.00000682, -8)
	proceed := decimal.NewFromFloatWithExponent(0.00272829, -8)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, Order{
		OrderID:      "55eb2c82-4184-4a24-8b6e-ee154b2f7eaf",
		MarketSymbol: "XRP-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     &quantity,
		Limit:        &limit,
		TimeInForce:  "GOOD_TIL_CANCELLED",
		FillQuantity: &fillQuantity,
		Commission:   &commission,
		Proceeds:     &proceed,
		Status:       "CLOSED",
		CreatedAt:    "2017-10-20T18:27:20.747Z",
		UpdatedAt:    "2017-10-20T18:27:20.763Z",
		ClosedAt:     "2017-10-20T18:27:20.763Z",
	})
}
func (this *BittrexAPIFixture) TestGetOrders() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	result, err := bittrex.GetOrders("open")
	quantity := decimal.NewFromFloatWithExponent(77.53046131, -8)
	limit := decimal.NewFromFloatWithExponent(0.00003528, -8)
	fillQuantity := decimal.NewFromFloatWithExponent(77.53046131, -8)
	commission := decimal.NewFromFloatWithExponent(0.00000682, -8)
	proceed := decimal.NewFromFloatWithExponent(0.00272829, -8)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, []Order{{
		OrderID:      "55eb2c82-4184-4a24-8b6e-ee154b2f7eaf",
		MarketSymbol: "XRP-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     &quantity,
		Limit:        &limit,
		TimeInForce:  "GOOD_TIL_CANCELLED",
		FillQuantity: &fillQuantity,
		Commission:   &commission,
		Proceeds:     &proceed,
		Status:       "OPEN",
		CreatedAt:    "2017-10-20T18:27:20.747Z",
		UpdatedAt:    "2017-10-20T18:27:20.763Z",
		ClosedAt:     "2017-10-20T18:27:20.763Z",
	}})

	result, err = bittrex.GetOrders("closed")
	this.So(err, should.BeNil)
	this.So(result[0].Status, should.Equal, "CLOSED")
}

func (this *BittrexAPIFixture) TestCreateOrder() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	quantity := decimal.NewFromFloat(5)
	limit := decimal.NewFromFloat(0.00039561)
	quantityWithExp := decimal.NewFromFloatWithExponent(5, 0)
	limitWithExp := decimal.NewFromFloatWithExponent(0.00039561, -8)
	order := Order{
		OrderID:      "",
		MarketSymbol: "ETH-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		TimeInForce:  "GOOD_TIL_CANCELLED",
		Quantity:     &quantity,
		Limit:        &limit,
	}
	result, err := bittrex.CreateOrder(order)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, &Order{
		OrderID:      "fab677a0-510e-456e-b450-8a75cea69f5d",
		MarketSymbol: "ETH-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     &quantityWithExp,
		Limit:        &limitWithExp,
		TimeInForce:  "GOOD_TIL_CANCELLED",
		Status:       "OPEN",
		CreatedAt:    "2020-09-08T05:08:40.84Z",
		UpdatedAt:    "2020-09-08T05:08:40.84Z",
		ClosedAt:     "",
	})
}

func (this *BittrexAPIFixture) TestCreateOrderError() {
	client := &fakeBittrexClient{EnableErrors: true}
	quantity := decimal.NewFromFloat(5)
	limit := decimal.NewFromFloat(0.00039561)
	bittrex := NewBittrexAPI(client, "")
	order := Order{
		OrderID:      "",
		MarketSymbol: "ETH-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		TimeInForce:  "GOOD_TIL_CANCELLED",
		Quantity:     &quantity,
		Limit:        &limit,
	}
	minTradeError := "MIN_TRADE_REQUIREMENT_NOT_MET"
	result, err := bittrex.CreateOrder(order)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, &Order{
		Quantity:     nil,
		Limit:        nil,
		Code:         &minTradeError,
	})
}

func (this *BittrexAPIFixture) TestCancelOrder() {
	client := &fakeBittrexClient{}
	bittrex := NewBittrexAPI(client, "")
	orderId := "fab677a0-510e-456e-b450-8a75cea69f5d"
	quantityWithExp := decimal.NewFromFloatWithExponent(5, 0)
	limitWithExp := decimal.NewFromFloatWithExponent(0.00039561, -8)
	result, err := bittrex.CancelOrder(orderId)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, &Order{
		OrderID:      orderId,
		MarketSymbol: "ETH-BTC",
		Direction:    "BUY",
		OrderType:    "LIMIT",
		Quantity:     &quantityWithExp,
		Limit:        &limitWithExp,
		TimeInForce:  "GOOD_TIL_CANCELLED",
		Status:       "OPEN",
		CreatedAt:    "2020-09-08T05:08:40.84Z",
		UpdatedAt:    "2020-09-08T05:08:40.84Z",
		ClosedAt:     "",
	})
}

func (this *BittrexAPIFixture) TestCancelOrderError() {
	client := &fakeBittrexClient{EnableErrors: true}
	bittrex := NewBittrexAPI(client, "")
	orderId := "fab677a0-510e-456e-b450-8a75cea69f5d"
	orderNotOpen := "ORDER_NOT_OPEN"
	result, err := bittrex.CancelOrder(orderId)
	this.So(err, should.BeNil)
	this.So(result, should.Resemble, &Order{
		Quantity:     nil,
		Limit:        nil,
		Code:         &orderNotOpen,
	})
}

///////////////////////////////////////

type fakeBittrexClient struct{
	EnableErrors bool
}

func (this *fakeBittrexClient) Do(method, uri, payload string, authenticate bool) ([]byte, error) {
	switch uri {
	case "/balances":
		return []byte("[{\"currencySymbol\": \"BTC\",\"total\": \"0.00000000\",\"available\": \"0.00000000\",\"updatedAt\": \"2019-10-29T20:25:10.16Z\"},{\"currencySymbol\": \"LTC\",\"total\": \"0\",\"available\": \"0\",\"updatedAt\": \"2020-09-03T21:27:53.8210894Z\"}]"), nil
	case "/currencies/fakesymbol":
		return []byte("{}"), nil
	case "/markets/fakesymbol":
		return []byte("{\"symbol\":\"ETH-BTC\",\"baseCurrencySymbol\":\"ETH\",\"quoteCurrencySymbol\":\"BTC\",\"minTradeSize\":\"0.01000000\",\"precision\":8,\"status\":\"ONLINE\",\"createdAt\":\"2015-08-14T09:02:24.817Z\",\"notice\":\"\",\"prohibitedIn\":[],\"associatedTermsOfService\":[]}"), nil
	case "/markets":
		return []byte("[{\"symbol\": \"4ART-BTC\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"BTC\",\"minTradeSize\": \"10.00000000\",\"precision\": 8,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:29.833Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]},{\"symbol\": \"4ART-USDT\",\"baseCurrencySymbol\": \"4ART\",\"quoteCurrencySymbol\": \"USDT\",\"minTradeSize\": \"10.00000000\",\"precision\": 5,\"status\": \"ONLINE\",\"createdAt\": \"2020-06-10T15:05:40.98Z\",\"notice\": \"\",\"prohibitedIn\": [\"US\"]}]"), nil
	case "/markets/fakesymbol/summary":
		return []byte("{\"symbol\":\"ETH-BTC\",\"high\":\"0.03894964\",\"low\":\"0.03650000\",\"volume\":\"18494.04035144\",\"quoteVolume\":\"696.42899671\",\"percentChange\":\"-3.33\",\"updatedAt\":\"2020-09-04T04:37:45.107Z\"}"), nil
	case "/markets/summaries":
		return []byte("[{\"symbol\": \"4ART-BTC\",\"high\": \"0.00000275\",\"low\": \"0.00000249\",\"volume\": \"54499.59344453\",\"quoteVolume\": \"0.13917073\",\"percentChange\": \"10.44\",\"updatedAt\": \"2020-09-04T04:58:55.447Z\"},{\"symbol\": \"4ART-USDT\",\"high\": \"0.02880000\",\"low\": \"0.02667000\",\"volume\": \"48259.53706735\",\"quoteVolume\": \"1320.75839607\",\"percentChange\": \"-6.11\",\"updatedAt\": \"2020-09-04T04:33:20.01Z\"}]"), nil
	case "/markets/fakesymbol/ticker":
		return []byte("{\"symbol\":\"ETH-BTC\",\"lastTradeRate\":\"0.03760069\",\"bidRate\":\"0.03760103\",\"askRate\":\"0.03762798\"}"), nil
	case "/markets/tickers":
		return []byte("[{\"symbol\": \"ETH-BTC\",\"lastTradeRate\": \"0.03760069\",\"bidRate\": \"0.03760103\",\"askRate\": \"0.03762798\"},{\"symbol\": \"ETH-FAKE\",\"lastTradeRate\": \"1.03760069\",\"bidRate\": \"1.03760103\",\"askRate\": \"1.03762798\"}]"), nil
	case "/orders/fakeOrder":
		return []byte("{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"CLOSED\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}"), nil
	case "/orders/open":
		return []byte("[{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"OPEN\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}]"), nil
	case "/orders/closed":
		return []byte("[{\"id\": \"55eb2c82-4184-4a24-8b6e-ee154b2f7eaf\",\"marketSymbol\": \"XRP-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"77.53046131\",\"limit\": \"0.00003528\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"fillQuantity\": \"77.53046131\",\"commission\": \"0.00000682\",\"proceeds\": \"0.00272829\",\"status\": \"CLOSED\",\"createdAt\": \"2017-10-20T18:27:20.747Z\",\"updatedAt\": \"2017-10-20T18:27:20.763Z\",\"closedAt\": \"2017-10-20T18:27:20.763Z\"}]"), nil
	case "/orders":
		if this.EnableErrors {
			return []byte("{\"code\": \"MIN_TRADE_REQUIREMENT_NOT_MET\"}"), nil
		}
		return []byte("{\"id\": \"fab677a0-510e-456e-b450-8a75cea69f5d\",\"marketSymbol\": \"ETH-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"5\",\"limit\": \"0.00039561\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"status\": \"OPEN\",\"createdAt\": \"2020-09-08T05:08:40.84Z\",\"updatedAt\": \"2020-09-08T05:08:40.84Z\"}"), nil
	case "/orders/fab677a0-510e-456e-b450-8a75cea69f5d":
		if this.EnableErrors {
			return []byte("{\"code\": \"ORDER_NOT_OPEN\"}"), nil
		}
		return []byte("{\"id\": \"fab677a0-510e-456e-b450-8a75cea69f5d\",\"marketSymbol\": \"ETH-BTC\",\"direction\": \"BUY\",\"type\": \"LIMIT\",\"quantity\": \"5\",\"limit\": \"0.00039561\",\"timeInForce\": \"GOOD_TIL_CANCELLED\",\"status\": \"OPEN\",\"createdAt\": \"2020-09-08T05:08:40.84Z\",\"updatedAt\": \"2020-09-08T05:08:40.84Z\"}"), nil
	}

	return nil, errors.New("test resource not found")
}

func (this *fakeBittrexClient) authenticate(request *http.Request, payload string, uri string, method string) error {
	return nil
}
