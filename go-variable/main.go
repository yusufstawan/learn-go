package main

import "fmt"

func main() {
	// deklarasi variabel dengan tipe data
	var firstName string = "John"
	var lastName string
	lastName = "wick"

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName+"!")

	// deklarasi variabel tanpa tipe data
	var firstName2 = "John"
	lastName2 := "wick"

	fmt.Printf("halo %s %s!\n", firstName2, lastName2)

	// Deklarasi Variabel Menggunakan Keyword new
	name := new(string)

	fmt.Println(name)
	fmt.Println(*name)
}
