package main

import "fmt"

func main() {
	const USDtoEUR = 0.849
	const USDtoRUB = 83.17
	const EURtoRUB = USDtoRUB / USDtoEUR

	amount := inputAmount()
	fmt.Printf("Вы ввели  %f", amount)
}

func inputAmount() float64 {
	fmt.Printf("Введите сумму")
	var amount float64
	fmt.Scan(&amount)

	return amount
}

func convertAmount(amount float64, currencyFrom string, currencyTo string) float64 {
	var convertedAmount float64

	return convertedAmount
}
