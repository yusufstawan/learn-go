package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// deklarasi dengan var
	var text = "Hello World"
	fmt.Println(text)

	text = "Hello World diubah"
	fmt.Println(text)

	// deklarasi dengan var dan tipe data
	var text1 string
	text1 = "Ini text 1"
	fmt.Println(text1)

	var text2 string = "Ini text 2"
	text2 = "Ini text 2 diubah"
	fmt.Println(text2)

	// deklarasi dengan menggunakan perantara ":="
	text3 := "Ini text 1"
	text3 = "ini text 1 diubah"
	fmt.Println(text3)

	text4 := "Ini text 2"
	text4 = "ini text 2 diubah"
	fmt.Println(text4)

	// constat
	const judul = "Ini judul"
	fmt.Println(judul)
	// tidak bisa dilakukan asign ulang
	// judul = "Ini judul diubah"
	// fmt.Println(judul)

	// numerik non-desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -12345655

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// tipe data boolean
	var exist bool = true
	exist = false
	fmt.Printf("exist: %t\n", exist)

	// tipe data string
	var message string = "Halo"
	fmt.Printf("Message: %s\n", message)
	// menggunakan backtick `
	var message1 = `Nama saya "John Doe".
	Salam kenal
	Mari belajar Golang.`

	fmt.Println(message1)

	var name = "John Doe"
	var age = "28"
	var sentence = `Halo nama saya "` + name + `" dan berumur "` + age + `" tahun.`
	fmt.Println(sentence)

	// konvensi data menggunakan teknik casting
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)

	var str = "Halo"
	var c string = string(str[0])
	fmt.Println(c)

	// package string
	var index1 = strings.Index("ethan hunt", "n")
	fmt.Println(index1)

	var textR = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(textR, find, replaceWith, 1)
	fmt.Println(newText1)

	var newText2 = strings.Replace(textR, find, replaceWith, 2)
	fmt.Println(newText2)

	var newText3 = strings.Replace(textR, find, replaceWith, -1)
	fmt.Println(newText3)

	var str1 = strings.Repeat("na", 4)
	fmt.Println(str1)

	var str2 = strings.ToLower("aLAy")
	fmt.Println(str2)

	var str3 = strings.ToUpper("eat!")
	fmt.Println(str3)

	// package strconv
	var str4 = "123"
	var num, err = strconv.Atoi(str4)

	if err == nil {
		fmt.Println(num)
		fmt.Println(err)
	}

	var nump = 124
	fmt.Println(strconv.Itoa(nump))

	// ParseInt()
	var str5 = "123"
	fmt.Println(strconv.ParseInt(str5, 10, 64))

	var strx = "false"
	var bul, _ = strconv.ParseBool(strx)

	fmt.Println(bul) // true

	var bulx = true
	var strxx = strconv.FormatBool(bulx)

	fmt.Println(strxx) // true
}
