package converter

//Exchanger exchanges the amount passed in based on the pair
func Exchanger(amount float64, pair string) (float64, error) {
	switch pair {
	case "KSH/NAIRA":
		conv := amount / 0.045
		return conv, nil
	case "KSH/GHS":
		conv := amount / 50.00
		return conv, nil
	case "GHS/KSH":
		conv := amount / 0.20
		return conv, nil
	case "GHS/NAIRA":
		conv := amount / 0.075
		return conv, nil
	case "NAIRA/KSH":
		conv := amount / 80.05
		return conv, nil
	case "NAIRA/GHS":
		conv := amount / 105.00
		return conv, nil
	}

	return 0, nil
}
