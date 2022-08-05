package main

import "fmt"

func main() {
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "Ruby"
	// languages[2] = "Javascript"
	// languages[3] = "C#"
	// languages[4] = "Python"

	languages := [...]string{"Go", "Ruby", "Javascript", "C#", "Python", "Kotlin"}

	for index, lang := range languages {
		fmt.Println("Index :", index, "Language :", lang)
	}

	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)
}
