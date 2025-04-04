package main

import "fmt"

func main() {
	var accountBalance = 1000.00

	fmt.Println("Welcome to the GO BANK!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Println("Your choice: ")
	fmt.Scan(&choice)

	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		if (withdrawalAmount > accountBalance) {
			fmt.Println("You do not have sufficient balance!")
		} else {
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount:", accountBalance)
		}
	} else {
		fmt.Println("Goodbye!")
	}
}