// Profit calculator
package main

import (
	"fmt"
)

func main() {
	//Tax rate

	const TaxRate = 0.25

	//Revenue

	var Revenue float64

	//Expensses

	var Expenses float64

	//Read user input
	fmt.Println("Enter revenue:")
	fmt.Scan(&Revenue)
	fmt.Println("Enter expenses:")
	fmt.Scan(&Expenses)

	//Earnings EBT

	var EBT = Revenue - Expenses

	//Profit EET

	var EET = (Revenue - Expenses) * (1 - TaxRate)

	//Ratio EBT /Ratio

	var EBTRatio = (EET / EBT) * 100

	//Output EBT EET /Ratio

	fmt.Println("Earning before tax:", EBT)
	fmt.Println("Earning after tax:", EET)
	fmt.Println("EBT/EET ratio:", EBTRatio, "%")
}
