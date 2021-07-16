package models

import "math"

type Currency struct {
	Name         string  `json:"name"`
	CurrencyPair string  `json:"currency_pair"`
	Amount       float64 `json:"amount"`
}

type CurrencyConvert struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

//Exchanger exchanges the amount passed in based on the pair
// exchange rates are based on (2021-07-16 @11:30am) rates
func Exchanger(amount float64, pair string) (float64, error) {
	//Round to nearest 2 decimal places
	amount = math.Round(amount*100) / 100

	switch pair {
	case "KSH/NGN":
		conv := amount / 0.26
		return conv, nil
	case "KSH/GHS":
		conv := amount / 18.17
		return conv, nil
	case "GHS/KSH":
		conv := amount / 0.055
		return conv, nil
	case "GHS/NGN":
		conv := amount / 0.014
		return conv, nil
	case "NGN/KSH":
		conv := amount / 3.80
		return conv, nil
	case "NGN/GHS":
		conv := amount / 69.13
		return conv, nil
	}

	return 0, nil
}
