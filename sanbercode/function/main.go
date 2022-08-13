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

// Variadic Function
func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Function Dengan Parameter Biasa & Variadic
func yourHobbies(name string, hobbies ...string) {
	fmt.Println("Hello, my name is ", name)
	fmt.Println("My hobbies are: ")
	for _, hobby := range hobbies {
		fmt.Println(hobby)
	}
}

// Function as Parameter
// function yang menggunakan function sebagai parameternya
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

// function yang digunakan sebagai parameter
func spamFilter(name string) string {
	if name == "Kasar" {
		return "..."
	} else {
		return name
	}
}

func main() {
	printHello()
	printAngka(10, 20)

	// langsung panggil dalam print
	fmt.Println(introduction("Yusuf"))

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

	// variadic function
	var total = sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)

	// Pengisian Parameter Fungsi Variadic Menggunakan Data Slice
	var numbers = []int{2, 6, 7, 8, 9, 10}
	var total2 = sum(numbers...)
	fmt.Println(total2)

	yourHobbies("John", "Gaming", "Reading", "Coding")

	var hobbies = []string{"Sleeping", "Eating"}
	yourHobbies("Yusuf", hobbies...)

	// Closure Function
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers3 = []int{2, 6, 7, 8, 9, 10}
	var min, max = getMinMax(numbers3)
	fmt.Println("data :", numbers)
	fmt.Println("min :", min)
	fmt.Println("max :", max)

	// Function as Parameter
	sayHelloWithFilter("Yusuf", spamFilter)
	sayHelloWithFilter("Kasar", spamFilter)
}
