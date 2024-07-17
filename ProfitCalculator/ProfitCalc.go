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
	ratio := ebt / profit
	fmt.Printf("ETB: %v\nRatio: %v\nProfit: %v\n", ebt, ratio, profit)
}
