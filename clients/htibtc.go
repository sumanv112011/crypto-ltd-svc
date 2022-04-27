package clients

import (
	"encoding/json"
	"fmt"
	"net/http"

	"goprojects/crypto-ltd-svc/config"
	"goprojects/crypto-ltd-svc/models"
)

// GetTicker will get the price details of symbol
func (c *client) GetTickers() (map[string]models.Ticker, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v%v", config.HITBTCURL, config.TickerURI), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpc.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp map[string]models.Ticker

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, err
	}

	return resp, nil
}
