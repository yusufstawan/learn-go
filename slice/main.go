package main

import "fmt"

func main() {
	var fruits = []string{"apple", "banana", "orange", "grape", "strawberry"}
	var newFruits = fruits[0:3]

	fmt.Println(newFruits)
}
