package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("TaxRate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formatOutput(ebt)
	formatOutput(profit)
	formatOutput(ratio)
}

func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate/100))
	ratio := ebt / profit
	
	return ebt, profit, ratio
}

func formatOutput(value float64){
	fmt.Printf("%.2f\n", value)
}