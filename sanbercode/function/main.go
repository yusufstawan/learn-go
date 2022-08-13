package main

import "fmt"

func printHello() {
	fmt.Println("Hello")
}

func printAngka(angka1 int, angka2 int) {
	fmt.Println("angka pertama", angka1)
	fmt.Println("angka kedua", angka2)
}

// return single value
func introduction(name string) string {
	return "Hello my name is " + name
}

// return multiple value
func introduction2(firstName string, lastName string) (string, string) {
	introFirstName := "Hello my first name is " + firstName
	introLastName := "Hello my last name is " + lastName
	return introFirstName, introLastName
}

// Function Predefined return value
func tambahAngka(firstNumber int, lastNumber int) (jumlah int) {
	jumlah = firstNumber + lastNumber
	return
}

func tampilkanKatadanAngka() (firstWord, secondWord string, number int) {
	firstWord = "Halo"
	secondWord = "Dunia"
	number = 10
	return
}

func main() {
	printHello()
	printAngka(10, 20)

	// langsung panggil dalam print
	fmt.Println(introduction(("Yusuf")))

	// menggunakan variable
	result := introduction("Yusuf")
	fmt.Println(result)

	// function as value
	secondResult := introduction
	fmt.Println(secondResult("Yusuf"))

	// menggunakan 2 variable
	firstName, lastName := introduction2("Yusuf", "Setiyawan")
	fmt.Println(firstName, lastName)

	// menggunakan 2 variable satu tidak dogunakan
	firstName2, _ := introduction2("Yusuf", "Setiyawan")
	fmt.Println(firstName2)

	hasil := tambahAngka(3, 2)
	fmt.Println(hasil)

	kataPertama, kataKedua, angka := tampilkanKatadanAngka()
	fmt.Println(kataPertama, kataKedua, angka)
}
