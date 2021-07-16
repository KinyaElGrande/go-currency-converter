package models

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
func Exchanger(amount float64, pair string) (float64, error) {
	switch pair {
	case "KSH/NGN":
		conv := amount / 0.045
		return conv, nil
	case "KSH/GHS":
		conv := amount / 50.00
		return conv, nil
	case "GHS/KSH":
		conv := amount / 0.20
		return conv, nil
	case "GHS/NGN":
		conv := amount / 0.075
		return conv, nil
	case "NGN/KSH":
		conv := amount / 80.05
		return conv, nil
	case "NGN/GHS":
		conv := amount / 105.00
		return conv, nil
	}

	return 0, nil
}
