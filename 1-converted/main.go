package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	EUR = "EUR"
	USD = "USD"
	RUB = "RUB"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	// Курс к USD: сколько USD в одной единице валюты
	var toUSD = map[string]float64{
		USD: 1.0,         // 1 USD = 1 USD
		EUR: 1.0 / 0.849, // 1 EUR = 1/0.849 USD
		RUB: 1.0 / 83.17, // 1 RUB = 1/83.17 USD
	}

	currencyFrom := inputCurrency("Введите исходную валюту. Доступные: EUR, USD, RUB", &toUSD)
	fmt.Printf("Вы ввели: %s\n", currencyFrom)

	amount := inputAmount()
	fmt.Printf("Вы ввели: %.2f\n", amount)

	currencyTo := inputCurrencyTo(currencyFrom, &toUSD)

	convertedAmount := convertAmount(amount, currencyFrom, currencyTo, &toUSD)

	fmt.Printf("Результат: %s\n", convertedAmount)
}

func inputAmount() float64 {
	var amount float64
	for {
		fmt.Print("Введите сумму: ")
		_, _ = fmt.Scan(&amount)
		if amount > 0 {
			return amount
		}
		fmt.Println("Сумма должна быть больше нуля.")
	}
}

func inputCurrency(prompt string, toUSDPointer *map[string]float64) string {
	for {
		fmt.Println(prompt)
		text, _ := reader.ReadString('\n')
		cur := strings.ToUpper(strings.TrimSpace(text))

		if checkCurrency(cur, toUSDPointer) {
			return cur
		}
		// Сообщение уже печатается в checkCurrency
	}
}

func inputCurrencyTo(currencyFrom string, toUSDPointer *map[string]float64) string {
	// Список доступных валют формируем из ключей map, исключая исходную
	available := make([]string, 0, len(*toUSDPointer)-1)
	for cur := range *toUSDPointer {
		if cur != currencyFrom {
			available = append(available, cur)
		}
	}
	for {
		fmt.Printf("Введите целевую валюту. Доступные: %s\n", strings.Join(available, ", "))
		text, _ := reader.ReadString('\n')
		cur := strings.ToUpper(strings.TrimSpace(text))

		if !checkCurrency(cur, toUSDPointer) {
			continue
		}
		if currencyFrom == cur {
			fmt.Printf("Вы ввели некорректную валюту: %s (совпадает с исходной)\n", cur)
			continue
		}
		return cur
	}
}

// Конвертация через базу USD: amount * toUSD[from] / toUSD[to]
func convertAmount(amount float64, currencyFrom, currencyTo string, toUSDPointer *map[string]float64) string {
	usd := amount * (*toUSDPointer)[currencyFrom]
	target := usd / (*toUSDPointer)[currencyTo]
	return fmt.Sprintf("%.2f", target)
}

func checkCurrency(cur string, toUSDPointer *map[string]float64) bool {
	if _, ok := (*toUSDPointer)[cur]; ok {
		return true
	}
	fmt.Printf("Вы ввели некорректную валюту: %s\n", cur)
	return false
}
