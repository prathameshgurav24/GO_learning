package main

import (
	"bank/fileops" //user define package
	"fmt"

	"github.com/Pallinder/go-randomdata" //third-party package
)

const fileName = "balance.txt"

func main() {

	fmt.Print(randomdata.PhoneNumber())

	currentBalance, err := fileops.GetFloatFromFile(fileName)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Cant Continue, Sorry.")
	}
	var amount float64
	var choice int

	for {
		presentOptions()

		fmt.Print("Your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("currentBalance is: ", currentBalance)
		case 2:
			fmt.Println("Enter amount to be deposited: ")
			fmt.Scan(&amount)
			currentBalance += amount
			fmt.Printf("Amount of %.2f is deposited and current balance is %.2f\n", amount, currentBalance)
		case 3:
			fmt.Println("Enter amount to be withdrawn: ")
			fmt.Scan(&amount)
			if amount <= currentBalance {
				currentBalance -= amount
			} else {
				fmt.Println("Invalid Amount. Current balance is too low XXX")
			}
			fmt.Printf("Amount of %.2f is withdrawn and current balance is %.2f\n", amount, currentBalance)
		case 4:
			fileops.WriteFloatToFile(currentBalance, fileName)
			fmt.Println("Thank you...Goodbye")
			return
		default:
			fmt.Println("Enter valid option")
		}
	}

}
