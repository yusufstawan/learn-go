package main

import "fmt"

func main() {
	// Tipe Data Numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// Tipe Data Numerik Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.10f\n", decimalNumber)

	// Tipe Data bool (Boolean)
	var exist bool = true

	fmt.Printf("exist? %t \n", exist)

	// Tipe Data string
	var message string = "Hello World"
	fmt.Printf("message: %s\n", message)

	var pesan = `Nama saya "John Wick".
	Salam kenal.
	Mari belajar "Golang".`

	fmt.Println(pesan)

	// Nilai nil & Zero Value
}
