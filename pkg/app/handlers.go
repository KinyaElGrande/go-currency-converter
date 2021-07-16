package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"currecy-converter/pkg/models"
)

func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(fmt.Sprintf("content-type should be 'application/json', but not '%s'", ct)))
		return
	}

	var currencyConvert models.CurrencyConvert
	err = json.Unmarshal(bodyBytes, &currencyConvert)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var currency models.Currency
	currency.Name = currencyConvert.To
	currency.CurrencyPair = strings.ToUpper(fmt.Sprintf("%s/%s", currencyConvert.From, currencyConvert.To))

	exchanged, err := models.Exchanger(currencyConvert.Amount, currency.CurrencyPair)
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
