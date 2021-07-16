package main

import (
	"currecy-converter/converter"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Currency struct {
	Name         string `json:"name"`
	CurrencyPair string `json:"currency_pair"`
	Amount       float64 `json:"amount"`
}

type CurrencyConvert struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func convertCurrency(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var currencyConvert CurrencyConvert
	err = json.Unmarshal(bodyBytes, &currencyConvert)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var currency Currency
	currency.Name = currencyConvert.To
	currency.CurrencyPair = strings.ToUpper(fmt.Sprintf("%s/%s",currencyConvert.From, currencyConvert.To))

	exchanged, err := converter.Exchanger(currencyConvert.Amount, currency.CurrencyPair)
	if err != nil {
		return 
	}
	currency.Amount = exchanged
	
	//response body 
	jsonBytes, err := json.Marshal(currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/convert", convertCurrency)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
