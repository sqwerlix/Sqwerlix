// Salary calculator
package main

import "fmt"

func main() {
	//Income tax rate swe
	const MunicipalTaxRate = 0.3238
	const CemetryTaxRate = 0.0028
	const EmploymentTaxRate = 0.239
	const TaxDeduction = 0.07

	var salary float64

	//Print statement to get user input.
	fmt.Println("Enter salary:")
	fmt.Scan(&salary, "SEK")

	//Calculate tax
	municipalTax := salary * MunicipalTaxRate
	cemetryTax := salary * CemetryTaxRate
	employmentTax := salary * EmploymentTaxRate
	tax := municipalTax + cemetryTax + employmentTax - TaxDeduction
	netSalary := salary*(EmploymentTaxRate+CemetryTaxRate+TaxDeduction) + salary

	//Calculate net salary
	fmt.Println("Net salary:", netSalary, "SEK")
	fmt.Println("Total tax amount:", tax, "SEK")
	fmt.Println("Net salary without employment tax:", salary, "SEK")

}
