package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetSymbolPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	symbol := vars["symbol"]

	// if symbol is all, send request to other handler.
	if symbol == "all" {
		GetAllSymbolPrices(w, r)

		return
	}

	// get the price of symbol
	price, err := svc.GetSymbolPrice(symbol)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error: %v", err.Error())))

		return
	}

	// convert data to json format/bytes
	res, err := json.Marshal(price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err.Error())))

		return
	}

	// send response to user.
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllSymbolPrices(w http.ResponseWriter, r *http.Request) {
	// get the price of symbol
	prices, err := svc.GetAllSymbolPrice()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("Error: %v", err.Error())))
	}

	// convert data to json format/bytes
	res, err := json.Marshal(prices)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err.Error())))
	}

	// send response to user.
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
