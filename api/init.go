package api

import (
	"goprojects/crypto-ltd-svc/config"
	"goprojects/crypto-ltd-svc/service"

	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var svc service.Manager

func NewCryptoSVC() {
	// initialize service layer.
	svc = service.NewManager()

	// sync prices for ever 1 min in real-time into in-memory
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			fmt.Println("syncing prices with in-memory")

			// sync price with inmemory
			svc.SyncPrices()
		}
	}()

	fmt.Printf("Started Crypto svc on port  %v\n %v\n", config.AppPort, image())

	// initialize routes
	routes()
}

func routes() {
	r := mux.NewRouter()

	r.HandleFunc("/v1/currency/{symbol}", GetSymbolPrice).Methods(http.MethodGet)
	r.HandleFunc("/v1/currency/all", GetAllSymbolPrices).Methods(http.MethodGet)

	http.ListenAndServe(config.AppPort, r)
}

func image() string {
	return `
	_______  _______           _______ _________ _______    _______           _______ 
	(  ____ \(  ____ )|\     /|(  ____ )\__   __/(  ___  )  (  ____ \|\     /|(  ____ \
	| (    \/| (    )|( \   / )| (    )|   ) (   | (   ) |  | (    \/| )   ( || (    \/
	| |      | (____)| \ (_) / | (____)|   | |   | |   | |  | (_____ | |   | || |      
	| |      |     __)  \   /  |  _____)   | |   | |   | |  (_____  )( (   ) )| |      
	| |      | (\ (      ) (   | (         | |   | |   | |        ) | \ \_/ / | |      
	| (____/\| ) \ \__   | |   | )         | |   | (___) |  /\____) |  \   /  | (____/\
	(_______/|/   \__/   \_/   |/          )_(   (_______)  \_______)   \_/   (_______/
																					   

	`
}
