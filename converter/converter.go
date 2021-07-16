package converter

import "fmt"

// Currency contains ID and Description of an Currency.
type Currency struct {
	ID          string `json:"id"`
}

type convertCurrencyResponse struct {
	From       string  `json:"from"`
	To         string  `json:"to"`
	FromAmount float64 `json:"from_amount"`
	ToAmount   float64 `json:"to_amount"`
}

func ConvertCurrency(from, to Currency, amount float64) (float64, error) {
	currencyPair := fmt.Sprintf("%s/%s", from, to)

}