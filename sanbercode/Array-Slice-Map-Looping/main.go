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

	// var fruits = []string{"apple", "grape", "banana"}
	// var bFruits = fruits[0:2]

	// fmt.Println(cap(bFruits)) // 3
	// fmt.Println(len(bFruits)) // 2

	// fmt.Println(fruits)  // ["apple", "grape", "banana"]
	// fmt.Println(bFruits) // ["apple", "grape"]

	// var cFruits = append(bFruits, "papaya")

	// fmt.Println(fruits)  // ["apple", "grape", "papaya"]
	// fmt.Println(bFruits) // ["apple", "grape"]
	// fmt.Println(cFruits) // ["apple", "grape", "papaya"]

	// fungsi copy
	// dst := make([]string, 3)
	// src := []string{"watermelon", "pinnaple", "apple", "orange"}
	// n := copy(dst, src)

	// fmt.Println(dst) // watermelon pinnaple apple
	// fmt.Println(src) // watermelon pinnaple apple orange
	// fmt.Println(n)   // 3

	// ======================= MAP =======================
	// var chicken map[string]int
	// chicken = map[string]int{}

	// chicken["januari"] = 50
	// chicken["februari"] = 40

	// fmt.Println("januari", chicken["januari"]) // januari 50
	// fmt.Println("mei", chicken["mei"])         // mei 0

	// menghapus elemen map
	// var chicken = map[string]int{"januari": 50, "februari": 40}

	// fmt.Println(len(chicken)) // 2
	// fmt.Println(chicken)

	// delete(chicken, "januari")

	// fmt.Println(len(chicken)) // 1
	// fmt.Println(chicken)

	// var chicken = map[string]int{"januari": 50, "februari": 40}
	// var value, isExist = chicken["januari"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	// ======================= Looping =======================
	// Perulangan Menggunakan Keyword for

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Angka", i)
	// }

	// Penggunaan Keyword for Dengan Argumen Hanya Kondisi

	// var i = 0
	// for i < 5 {
	// 	fmt.Println("Angka", i)
	// 	i++
	// }

	// Penggunaan Keyword for Tanpa Argumen

	// var i = 0

	// for {
	// 	fmt.Println("Angka", i)

	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for range
	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}
}
