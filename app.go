package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, revenueErr := getUserInput("Revenue: ")
	if revenueErr != nil {
		fmt.Println("ERROR")
		fmt.Println(revenueErr)
		fmt.Println("----------")
		panic("Can't continue, sorry")
	}

	expenses, expensesErr := getUserInput("Expenses: ")
	if expensesErr != nil {
		fmt.Println("ERROR")
		fmt.Println(expensesErr)
		fmt.Println("----------")
		panic("Can't continue, sorry")
	}

	taxRate, taxRateErr := getUserInput("TaxRate: ")
	if taxRateErr != nil {
		fmt.Println("ERROR")
		fmt.Println(taxRateErr)
		fmt.Println("----------")
		panic("Can't continue, sorry")
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	formatOutput(ebt)
	formatOutput(profit)
	formatOutput(ratio)

	saveResultsToFile(ebt, profit, ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 1, errors.New("Invalid input.")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - (taxRate / 100))
	ratio := ebt / profit

	return ebt, profit, ratio
}

func formatOutput(value float64) {
	fmt.Printf("%.2f\n", value)
}

func saveResultsToFile(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.1f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
