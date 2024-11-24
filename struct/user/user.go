package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstname string
	lastname  string
	birthDate string
	createdAt time.Time
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("required all fields")
	}
	return &User{
		firstname: userFirstName,
		lastname:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstname, u.lastname, u.birthDate, u.createdAt)
}
func (u *User) ClearUserName() {
	u.firstname = ""
	u.lastname = ""
}
