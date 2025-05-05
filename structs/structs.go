package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName,
		lastName,
		birthdate,
		time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputUserDetails(&appUser)

}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
