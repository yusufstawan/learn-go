package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "John"
	user.LastName = "Doe"
	user.Email = "mail@gmail.com"
	user.IsActive = true

	user2 := User{1, "Yusuf", "Setiyawan", "yusuf@gmail.com", true}

	displayUser1 := displayUser(user)

	fmt.Println(displayUser1)
	fmt.Println(user2)
}

func displayUser(user User) string {
	result := fmt.Sprintf("Name: %s %s\nEmail: %s\nIs Active: %t", user.FirstName, user.LastName, user.Email, user.IsActive)
	return result
}
