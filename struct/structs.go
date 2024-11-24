package main

import (
	"fmt"
	"struct/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Print(err)
		return
	}

	// ... do something awesome with that gathered data!
	appUser.OutputUserDetails()
	// appUser.ClearUserName()
	// appUser.OutputUserDetails()

	var admin user.Admin
	admin, _ = user.NewAdmin("rohit@aus.com", "hitman200", *appUser)

	admin.OutputUserDetails()

	fmt.Print(admin)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
