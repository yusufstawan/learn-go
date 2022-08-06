package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	var total int
	for _, score := range scores {
		total += score
	}

	length := len(scores)
	average := float64(total) / float64(length)

	fmt.Println(average)
}
