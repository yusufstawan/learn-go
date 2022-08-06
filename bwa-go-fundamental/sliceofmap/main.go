package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Yusuf", "score": "A"},
		{"name": "Setiyawan", "score": "B"},
		{"name": "Rizki", "score": "C"},
	}

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}
