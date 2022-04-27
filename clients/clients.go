package clients

import (
	"goprojects/crypto-ltd-svc/models"
	"net/http"
	"time"
)

type Clients interface {
	GetTickers() (map[string]models.Ticker, error)
}

type client struct {
	httpc *http.Client
}

func NewClient() Clients {
	return &client{
		httpc: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
