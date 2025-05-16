package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	// receiver argument
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func (u *user) clearUsername() {
	u.firstName = ""
	u.lastName = ""
}

// returning a pointer instead of values
func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name & birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()
	appUser.clearUsername()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
