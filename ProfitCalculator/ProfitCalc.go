// Profit calculator
package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Enter revenue")
	fmt.Scan(&revenue)
	fmt.Println("Enter expenses")
	fmt.Scan(&expenses)
	fmt.Println("Enter tax rate")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := (profit / ebt) * 100
	percentage_profit := (profit / revenue) * 100
	fmt.Printf("ETB: %v\nRatio: %.2f%%\nProfit: %v\nPercentage Profit: %.2f%%", ebt, ratio, profit, percentage_profit)
}
