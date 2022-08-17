package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("halo", s.name)
}

func main() {
	var john = student{"john wick", 21}
	john.sayHello()
}
