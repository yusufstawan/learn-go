package main

import "fmt"

func main() {
	// Seleksi Kondisi Menggunakan Keyword if, else if, & else
	// var point = 8

	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// Variabel Temporary Pada if - else
	// var point = 8840.0

	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%s perfect!\n", percent, "%")
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%s good\n", percent, "%")
	// } else {
	// 	fmt.Printf("%.1f%s not bad\n", percent, "%")
	// }

	// Seleksi Kondisi Menggunakan Keyword switch - case
	// var point = 6
	// switch point {
	// case 8:
	// 	fmt.Println("perfect")
	// case 7, 6, 5, 4:
	// 	fmt.Println("awesome")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you can be better")
	// 	}
	// }

	// Switch Dengan Gaya if - else
	var point = 8

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}
}
