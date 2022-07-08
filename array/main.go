package main

import "fmt"

func main() {
	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	// // names[4] = "ez" // baris kode ini menghasilkan error

	// fmt.Println(names[0], names[1], names[2], names[3])

	// Inisialisasi Nilai Awal Array
	// var fruits = [4]string{"apple", "banana", "orange", "grape"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi element \t", fruits)

	//Inisialisasi Nilai Array Dengan Gaya Vertikal
	// var fruits [4]string

	// cara horizontal
	// fruits = [4]string{"apple", "banana", "orange", "grape"}

	// cara vertikal
	// fruits1 = [4]string{
	// 	"apple",
	// 	"banana",
	// 	"orange",
	// 	"grape",
	// }

	// Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	// var numbers = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("data array :", numbers)
	// fmt.Println("jumlah element :", len(numbers))

	//Array multidimensi
	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1 :", numbers1)
	// fmt.Println("numbers2 :", numbers2)

	// Perulangan Elemen Array Menggunakan Keyword for
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}
}
