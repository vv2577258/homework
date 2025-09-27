package main

import (
	"fmt"
)

const EUR = "EUR"
const USD = "USD"
const RUB = "RUB"

func main() {

	currencyFrom := inputCurrency()

	fmt.Printf("Вы ввели  %s\n", currencyFrom)

	amount := inputAmount()
	fmt.Printf("Вы ввели  %.2f\n", amount)

	currencyTo := inputCurrencyTo(currencyFrom)

	convertedAmount := convertAmount(amount, currencyFrom, currencyTo)

	fmt.Printf("Результат: %s\n", convertedAmount)
}

func inputAmount() float64 {
	var amount float64

	for {
		fmt.Printf("Введите сумму\n")
		fmt.Scan(&amount)

		if amount <= 0 {
			continue
		}

		break
	}

	return amount
}

func inputCurrency() string {
	var curCurrency string

	for {
		fmt.Printf("Введите исходную валюту. Доступные валюты: %s, %s, %s\n", EUR, USD, RUB)
		fmt.Scanln(&curCurrency)

		isValid := checkCurrency(curCurrency)

		if !isValid {

			continue
		}

		break
	}

	return curCurrency
}

func inputCurrencyTo(currencyFrom string) string {

	var avalibleCurrency string

	switch currencyFrom {
	case EUR:
		avalibleCurrency = fmt.Sprintf("%s, %s", USD, RUB)
	case USD:
		avalibleCurrency = fmt.Sprintf("%s, %s", EUR, RUB)
	case RUB:
		avalibleCurrency = fmt.Sprintf("%s, %s", USD, EUR)
	default:
		panic("AVAILIBLE_CURRENCY_ERROR")
	}

	var curCurrency string

	for {
		fmt.Printf("Введите целевую валюту. Доступные валюты: %s\n", avalibleCurrency)
		fmt.Scanln(&curCurrency)
		isValid := checkCurrency(curCurrency)

		if !isValid {

			continue
		}

		if currencyFrom == curCurrency {
			fmt.Println("Вы ввели некорректную валюту: %s", curCurrency)

			continue
		}

		break
	}

	return curCurrency
}

func convertAmount(amount float64, currencyFrom string, currencyTo string) string {

	USDMap := map[string]float64{
		"EUR": 0.849, "RUB": 83.17,
	}

	const USDtoEUR = 0.849
	const USDtoRUB = 83.17
	const EURtoRUB = USDtoRUB / USDtoEUR

	var convertedAmount float64

	switch {
	case currencyFrom == USD && currencyTo == EUR:
		convertedAmount = amount * USDMap["EUR"]
	case currencyFrom == USD && currencyTo == RUB:
		convertedAmount = amount * USDMap["RUB"]
	case currencyFrom == EUR && currencyTo == RUB:
		convertedAmount = amount * EURtoRUB
	case currencyFrom == EUR && currencyTo == USD:
		convertedAmount = amount / USDtoEUR
	case currencyFrom == RUB && currencyTo == USD:
		convertedAmount = amount / USDtoRUB
	case currencyFrom == RUB && currencyTo == EUR:
		convertedAmount = amount / EURtoRUB
	}

	normalizeAmount := fmt.Sprintf("%.2f", convertedAmount)

	return normalizeAmount
}

func checkCurrency(curCurrency string) bool {
	if curCurrency == EUR || curCurrency == USD || curCurrency == RUB {

		return true
	}

	fmt.Println("Вы ввели некорректную валюту: %s", curCurrency)
	return false
}
