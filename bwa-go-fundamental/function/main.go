package main

import "fmt"

func main() {
	sentence := printMyResult("Saya sedang")
	fmt.Println(sentence)
}

func printMyResult(sentence string) string {
	newSentence := sentence + " belajar golang"
	return newSentence
}
