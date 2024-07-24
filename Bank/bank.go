package main

import (
	"fmt"
	"os"
)

var accountBalance float64

func main() {
	//Initial balance
	fmt.Println("Welcome to the Bank!")
	fmt.Println("What would you like to do?")
	for {
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit")
		fmt.Println("3. Withdraw")
		fmt.Println("4. Check transactions")
		fmt.Println("5. Exit")

		var choice int
		fmt.Scan(&choice)
		fmt.Println("Your choice is", choice)

		//Adding a switch statement to replace the if-else if-else block
		switch choice {
		case 1:
			//Calling function checkBalance
			checkBalance(accountBalance)
		case 2:
			//Calling function deposit
			accountBalance = deposit(accountBalance)
			saveBalance(accountBalance)
		case 3:
			//Calling function withdraw
			accountBalance = withdraw(accountBalance)
			saveBalance(accountBalance)
		case 4:
			//Calling function saveTransactions
			checkTransactions()
		case 5:
			fmt.Println("Exiting the program...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}

}

//Functions to check balance, deposit and withdraw.

func checkBalance(accountBalance float64) {
	fmt.Println("Your balance is", accountBalance)
}

func deposit(accountBalance float64) float64 {
	var depositAmount float64
	fmt.Println("Enter amount to deposit")
	fmt.Scan(&depositAmount)
	accountBalance += depositAmount
	saveTransactions("Deposit", depositAmount)
	return accountBalance
}

func withdraw(accountBalance float64) float64 {
	var withdrawAmount float64
	fmt.Println("Enter amount to withdraw")
	fmt.Scan(&withdrawAmount)
	accountBalance -= withdrawAmount
	saveTransactions("Withdraw", withdrawAmount)
	return accountBalance
}

func saveTransactions(transactionType string, transactionAmount float64) {
	//Code to save transactions
	file, err := os.OpenFile("transactions.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s: %f\n", transactionType, transactionAmount))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func checkTransactions() {
	//Code to check transactions
	file, err := os.ReadFile("transactions.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Transactions:\n", string(file))
}

func saveBalance(accountBalance float64) {
	//Code to save balance
	file, err := os.OpenFile("balance.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%f\n", accountBalance))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
