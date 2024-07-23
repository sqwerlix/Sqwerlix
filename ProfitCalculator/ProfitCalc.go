// Profit calculator
package main

import (
	"fmt"
)

func main() {

	//fmt.Println("Enter revenue")
	revenue := readUserInput("Enter revenue: ")
	//fmt.Println("Enter expenses")
	expenses := readUserInput("Enter expenses: ")
	//fmt.Println("Enter tax rate")
	taxRate := readUserInput("Enter tax rate: ")

	ebt, profit, ratio, percentage_profit := calculateProfit(revenue, expenses, taxRate)
	fmt.Printf("ETB: %v\nProfit: %v\nRatio: %.2f%%\nPercentage profit: %.2f%%", ebt, profit, ratio, percentage_profit)
}

// Read user input function
func readUserInput(InfoText string) float64 {
	var userInput float64
	fmt.Print(InfoText)
	fmt.Scan(&userInput)
	return userInput
}

// Calculate profit function
func calculateProfit(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := (profit / ebt) * 100
	percentage_profit := (profit / revenue) * 100
	return ebt, profit, ratio, percentage_profit
}
