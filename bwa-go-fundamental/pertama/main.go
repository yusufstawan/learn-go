package main

// import package
import (
	"fmt"
	"pertama/calculation"
)

func main() {
	fmt.Println("Halo, belajar Golang")

	result := calculation.Add(10, 20)
	fmt.Println(result)
}
