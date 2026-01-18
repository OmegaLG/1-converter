package main

import "fmt"

func main() {
	const usdToEur = 0.86
	const usdToRub = 77.83
	const eurToRub = usdToRub / usdToEur
}

func readCurrencyConversionInput() (float64, string, string) {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Введите количество конвертируемой валюты: ")
	fmt.Scan(&amount)
	fmt.Print("Введите код конвертируемой валюты (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)
	fmt.Print("Введите код валюты, в которую нужно конвертировать (USD, EUR, RUB): ")
	fmt.Scan(&toCurrency)

	return amount, fromCurrency, toCurrency
}

func convertCurrency(amount float64, fromCurrency, toCurrency string) float64 {
	return 0.0
}
