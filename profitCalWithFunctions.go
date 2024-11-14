package main

import "fmt"

func main() {
	revenue := getUserInput("Revenue : ")
	expenses :=getUserInput("Expenses : ")
	taxRate := getUserInput("Tax Rate : ")

	ebt, pofit,ratio := calculateFinancials(revenue, expenses,taxRate)

	fmt.Printf("%.1f",ebt)
	fmt.Printf("%.1f",profit)
	fmt.Printf("%.3f",ratio)
}

func calculateFinancials(revenue,expenses,taxRate float64)(float64, float64, float64){
	ebt := revenue-expenses
	profit :=ebt * (1-taxRate/100)
	ratio := ebt/profit
	return ebt,profit,ratio
}

func getUserInput(infoTest string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
