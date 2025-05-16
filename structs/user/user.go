package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// struct embedding
type Admin struct {
	email string
	password string
	User User
}

func (u *User) OutputUserDetails() {
	// receiver argument
	fmt.Println(u.firstName, u.lastName, u.createdAt)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) Admin {
	return Admin {
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthDate: "------",
			createdAt: time.Now(),
		},
	}
}

// constructor: exporting a pointer instead of values
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name & birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
