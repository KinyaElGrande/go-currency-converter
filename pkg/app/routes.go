package app

import "net/http"

//Run serves the application
func Run(port string) {
	http.HandleFunc("/convert", ConvertCurrency)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}
