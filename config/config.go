package config

var (
	HITBTCURL      string
	TickerURI      string
	AllowedSymbols map[string][]string

	AppPort string
)

func init() {
	HITBTCURL = "https://api.hitbtc.com/api/3"
	TickerURI = "/public/ticker"

	AppPort = ":8080"

	AllowedSymbols = map[string][]string{
		"ETHBTC": {"Ethereum", "BTC"},
		"BTCUSD": {"Bitcoin", "USD"},
	}
}
