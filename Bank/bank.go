package main

import (
	"fmt"
	"os"
	"time"
)

//

func main() {
	//Initial balance
	var accountBalance float64

	// Read balance from file
	file, err := os.Open("balance.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fscanln(file, &accountBalance)
	if err != nil {
		fmt.Println("Error reading balance from file:", err)
		return
	}

	//Welcome message
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
			fmt.Println("Check balance")
			//Calling function checkBalance
			checkBalance()
		case 2:
			fmt.Println("Deposit")
			//Calling function deposit
			accountBalance = deposit(accountBalance)
			saveBalance(accountBalance)
		case 3:
			fmt.Println("Withdraw")
			//Calling function withdraw
			accountBalance = withdraw(accountBalance)
			saveBalance(accountBalance)
		case 4:
			fmt.Println("Check transactions")
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

func checkBalance() {
	//Code to check balance
	file, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Balance:\n", string(file))
}

func deposit(accountBalance float64) float64 {
	var depositAmount float64
	fmt.Println("Enter amount to deposit")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Deposit amount cannot be negative or zero")
		return accountBalance
	}
	accountBalance += depositAmount
	saveTransactions("Deposit", depositAmount)
	fmt.Println("You have deposited:", depositAmount)
	return accountBalance
}

func withdraw(accountBalance float64) float64 {
	var withdrawAmount float64
	fmt.Println("Enter amount to withdraw")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount")
		return accountBalance
	}
	if withdrawAmount > accountBalance {
		fmt.Println("Insufficient funds")
		return accountBalance
	}
	accountBalance -= withdrawAmount
	saveTransactions("Withdraw", withdrawAmount)
	fmt.Println("You have withdrawn:", withdrawAmount)
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

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(fmt.Sprintf("%s: %s - %f\n", currentTime, transactionType, transactionAmount))
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

/*
0: No permissions
1: Execute permission
2: Write permission
3: Execute and write permissions
4: Read permission
5: Execute and read permissions
6: Read and write permissions
7: All permissions (read, write, execute)
*/
