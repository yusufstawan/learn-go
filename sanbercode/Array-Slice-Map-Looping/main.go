package main

import "fmt"

func main() {
	// ======================= Array =======================
	// var names [4]string
	// names[0] = "John"
	// names[1] = "Doe"
	// names[2] = "Frank"
	// names[3] = "Jack"

	// fmt.Println(names[0], names[1], names[2], names[3])

	// Inisialisasi nilai awal array
	// var names = [4]string{"John", "Doe", "Frank", "Jack"}

	// fmt.Println(names[0], names[1], names[2], names[3])

	// fruits := [4]string{
	// 	"apple",
	// 	"orange",
	// 	"banana",
	// 	"grape",
	// }

	// fmt.Println(fruits)

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	// var numbers = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))

	// Array Multidimensi
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// ======================= Slice =======================
	// Inisialisasi Slice
	// var fruits = []string{"apple", "orange", "banana", "grape"}
	// fmt.Println(fruits[0])

	// fungsi leb()
	// var fruits = []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(len(fruits)) // 4

	// append()
	// var fruits = []string{"apple", "grape", "banana"}
	// var cfruits = append(fruits, "papaya")

	// fmt.Println(fruits)
	// fmt.Println(cfruits)

	var fruits = []string{"apple", "grape", "banana"}
	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	fmt.Println(bFruits) // ["apple", "grape"]
	fmt.Println(cFruits) // ["apple", "grape", "papaya"]
}
