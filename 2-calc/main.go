package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var validOperation = []string{"AVG", "SUM", "MED"}

func main() {

	curOperation, err := inputOperation()

	if err != nil {
		panic("Invalid operation")
	}

	numbersArr, numErr := inputNumbers()

	if numErr != nil {
		panic("Invalid input numbers")
	}

	operationResult, calcErr := calculate(curOperation, numbersArr)

	if calcErr != nil {
		panic("Invalid calculate")
	}

	fmt.Println("Ваша операция ", curOperation)
	fmt.Println("Ваш результат ", operationResult)
}

func inputOperation() (string, error) {
	fmt.Printf("Введите операцию: %s \n", validOperation)

	var curOperation string

	fmt.Scanln(&curOperation)

	normalizeCurOperation := strings.ToUpper(curOperation)
	found := slices.Contains(validOperation, normalizeCurOperation)

	if !found {
		return "", errors.New("invalid operation")
	}

	return normalizeCurOperation, nil
}

func inputNumbers() ([]int, error) {
	fmt.Println("Введите числа через запятую")

	var numbersArr []int
	var numbersStr string
	fmt.Scanln(&numbersStr)

	fmt.Println("Ваши числа ", numbersStr)

	curNumbers := strings.Split(numbersStr, ",")

	for _, value := range curNumbers {
		normalizeValue := strings.TrimSpace(value)
		fmt.Println(normalizeValue)
		curNum, err := strconv.Atoi(normalizeValue)

		if err != nil {
			return numbersArr, errors.New("invalid number")
		}
		numbersArr = append(numbersArr, curNum)

	}

	return numbersArr, nil
}

func calculate(operation string, arrNum []int) (float64, error) {

	var result float64
	var sum int

	switch operation {
	case "AVG":
		for _, number := range arrNum {
			sum += number
		}
		result = float64(sum) / float64(len(arrNum))
	case "SUM":
		for _, number := range arrNum {
			sum += number
		}
		result = float64(sum)
	case "MED":
		result = calcMedian(arrNum)
	default:
		return 0, errors.New("invalid operation in calculate")
	}

	return result, nil
}

func calcMedian(arrNum []int) float64 {
	dataCopy := make([]int, len(arrNum))
	copy(dataCopy, arrNum)

	sort.Ints(dataCopy)
	var median float64
	l := len(dataCopy)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		median = (float64(dataCopy[l/2-1] + dataCopy[l/2])) / 2
	} else {
		median = float64(dataCopy[l/2])
	}

	return median
}
