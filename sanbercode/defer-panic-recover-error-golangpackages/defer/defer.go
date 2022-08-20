package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Running application")
}

func main() {
	runApplication()
}
