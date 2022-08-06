package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	result, err := calculate(10, 2, "=")
	if err != nil {
		fmt.Println("Terjadi kesalahan :", err)
	} else {
		fmt.Println(result)
	}
}

func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}

	return total
}

func calculate(number int, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("unknown operation")
	}

	return result, errorResult
}
