// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var names = []string{"John", "Wick"}
// 	printMessage("Hallo", names)
// }

// func printMessage(message string, arr []string) {
// 	var nameString = strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// fungsi dengan return

package main

import "fmt"

func main() {
	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d divided by %d is %d\n", m, n, res)
}
