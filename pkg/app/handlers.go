package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"

	"currecy-converter/pkg/models"
)

//ConvertCurrency is a handler function that handles the conversion
func ConvertCurrency(w http.ResponseWriter, r *http.Request) {
	//Read the body Request
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	//Validate Content Type
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

	//Validate the 3 allowed currencies (NGN), (GHS) and (KSH)
	if validErrs := currencyConvert.Validate(); len(validErrs) > 0 {
		err := map[string]interface{}{"validationError": validErrs}
		w.Header().Set("Content-type", "applciation/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	//Start preparing response
	var currency models.Currency
	currency.Name = currencyConvert.To
	currency.CurrencyPair = strings.ToUpper(fmt.Sprintf("%s/%s", currencyConvert.From, currencyConvert.To))

	exchanged, err := models.Exchanger(currencyConvert.Amount, currency.CurrencyPair)
	if err != nil {
		return
	}

	//Round to nearest 2 decimal places
	currency.Amount = math.Round(exchanged*100) / 100

	//response body
	jsonBytes, err := json.Marshal(currency)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
