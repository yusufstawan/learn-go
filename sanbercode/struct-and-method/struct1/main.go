package main

import (
	"fmt"
)

func main() {
	// struct dengan deklarasi
	type student struct {
		name  string
		grade int
	}

	var john student
	john.name = "john doe"
	john.grade = 2

	fmt.Println("Name:", john.name)
	fmt.Println("Grade:", john.grade)

	// struct literals
	// cara pertama
	var doe = student{}
	doe.name = "wick"
	doe.grade = 3

	// cara kedua tetapi isinya harus berurutab
	var wick = student{"doek", 4}

	// cara ketiga dengan nama property tetapi tidak harus berurutan
	var jack = student{name: "jack", grade: 5}

	fmt.Println("Student 1:", doe.name)
	fmt.Println("Student 2:", wick.name)
	fmt.Println("Student 3:", jack.name)
}
