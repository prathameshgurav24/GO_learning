package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(fileName)
	if err == nil {
		balanceText := string(data)
		balance, err := strconv.ParseFloat(balanceText, 64)
		if err != nil {
			return 1000, errors.New("failed to parse stored balance value")
		} else {
			return balance, nil
		}
	} else {
		return 1000, errors.New("failed to find balance.txt")
	}

}

func main() {

	currentBalance, err := readBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		panic("Cant Continue, Sorry.")
	}
	var amount float64
	var choice int

	for true {
		fmt.Println("\n-----Welcome to Go Bank-----")
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
			writeBalanceToFile(currentBalance)
			fmt.Println("Thank you...Goodbye")
			return
		default:
			fmt.Println("Enter valid option")
		}
	}

}
