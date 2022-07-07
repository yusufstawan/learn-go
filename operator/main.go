package main

import "fmt"

func main() {
	// Operator Perbandingan
	var value = (((2+6)%3)*4 - 1) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	// Operator Logika
	var left = false
	var right = true

	var leftAndRight = left && right
	fmt.Printf("left && right = %t \n", leftAndRight)

	var leftOrRight = left || right
	fmt.Printf("left || right \t(%t) \n", leftOrRight)

	var leftReverse = !left
	fmt.Printf("!left \t\t(%t) \n", leftReverse)
}
