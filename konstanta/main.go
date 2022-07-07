package main

import "fmt"

func main() {
	const firstName string = "John"
	fmt.Print("halo ", firstName, "!\n")

	// type inference
	const lastName = "wick"
	fmt.Print("halo ", lastName, "!\n")
}
