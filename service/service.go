package service

import (
	"fmt"
	cryptocli "goprojects/crypto-ltd-svc/clients"
	"goprojects/crypto-ltd-svc/config"
	"goprojects/crypto-ltd-svc/models"
)

var (
	tickers map[string]models.Ticker
)

type Manager interface {
	GetSymbolPrice(symbol string) (models.Currency, error)
	GetAllSymbolPrice() ([]models.Currency, error)
	SyncPrices()
}

type manager struct {
	client cryptocli.Clients
}

func NewManager() Manager {
	c := cryptocli.NewClient()

	// load initial prices
	t, err := c.GetTickers()
	if err != nil {
		panic(err)
	}

	tickers = t

	return &manager{
		client: c,
	}
}

// SyncPrices will sync real-time prices to in-memory
func (m *manager) SyncPrices() {
	t, err := m.client.GetTickers()
	if err != nil {
		panic(err)
	}

	tickers = t
}

func (m *manager) GetSymbolPrice(symbol string) (models.Currency, error) {
	var (
		symbolDetails []string
		isValidSymbol bool
	)

	if symbol != "" {
		if symbolDetails, isValidSymbol = config.AllowedSymbols[symbol]; !isValidSymbol {
			return models.Currency{}, fmt.Errorf("currently allowed symbols: %v", config.AllowedSymbols)
		}
	}

	// get price from in memory
	price := tickers[symbol]

	// build symbol data
	c := models.Currency{
		ID:          symbol[:3],
		FullName:    symbolDetails[0],
		Ask:         price.Ask,
		Bid:         price.Bid,
		Last:        price.Last,
		Open:        price.Open,
		Low:         price.Low,
		High:        price.High,
		FeeCurrency: symbolDetails[1],
	}

	return c, nil
}

func (m *manager) GetAllSymbolPrice() ([]models.Currency, error) {
	res := make([]models.Currency, 0)

	for key, values := range config.AllowedSymbols {
		// get price from in memory
		price := tickers[key]

		// build symbol data
		c := models.Currency{
			ID:          key[:3],
			FullName:    values[0],
			Ask:         price.Ask,
			Bid:         price.Bid,
			Last:        price.Last,
			Open:        price.Open,
			Low:         price.Low,
			High:        price.High,
			FeeCurrency: values[1],
		}

		res = append(res, c)
	}

	return res, nil
}
