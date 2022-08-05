package main

import "fmt"

func main() {
	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"Javascript": "hype",
	}

	for key, value := range myMap {
		fmt.Println("Key :", key, "Value :", value)
	}

	fmt.Println("==========")

	delete(myMap, "ruby")

	for key, value := range myMap {
		fmt.Println("Key :", key, "Value :", value)
	}
}
