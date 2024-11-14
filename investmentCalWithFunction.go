package main

import "fmt"

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	//fmt.Print("Investment Amount : ")
	outputText("Investment Amount : ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate : ")
	outputText("Expected Return Rate : ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years : ")
	outputText("Years : ")
	fmt.Scan(&years)
   
	func outputText(text string){
		fmt.Print(text)
	}

	futureValue , futureRealValue = calculateFutureValues(investmentAmount,expectedReturnRate,years)



	func calculateFutureValues(investmentAmount , expectedReturnRate,years float64)(float64,float64){
		fv = investmentAmount * math.Pow(1+expectedReturnRate/100,years)
		rfv = fv/math.Pow(1+inflationRate/100,years)
		return
	}
}
