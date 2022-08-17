package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (values):", numberA)
	fmt.Println("numberA (address):", &numberA)

	fmt.Println("numberB (values):", *numberB)
	fmt.Println("numberB (address):", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (values):", numberA)
	fmt.Println("numberA (address):", &numberA)
	fmt.Println("numberB (values):", *numberB)
	fmt.Println("numberB (address):", numberB)

}
