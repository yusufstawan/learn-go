package main

import "fmt"

func main() {
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

	// Penggunaan Keyword break & continue
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// Perulangan Bersarang
	for i := 1; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
