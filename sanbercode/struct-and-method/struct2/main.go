package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	// contoh 1
	var john = student{}
	john.name = "john"
	john.age = 21
	john.grade = 2

	fmt.Println("name  :", john.name)
	fmt.Println("age   :", john.age)
	fmt.Println("age   :", john.person.age)
	fmt.Println("grade :", john.grade)

	// contoh 2
	var doeData = person{name: "doe", age: 21}
	var doe = student{person: doeData, grade: 2}

	fmt.Println("name  :", doe.name)
	fmt.Println("age   :", doe.age)
	fmt.Println("grade :", doe.grade)
}
