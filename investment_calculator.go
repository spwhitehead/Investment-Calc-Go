package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Investment Amount: ")
	printText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Number of Years: ")
	printText("Number of Years: ")
	fmt.Scan(&years)

	// fmt.Print("Expected Rate of Return: ")
	printText("Expected Rate of Return: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	
	// formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)


	// fmt.Println("Future Value: $", futureValue)
	///fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Printf("Future Value (adjusted for inflation): $%v\n", futureRealValue)

	// fmt.Print(formattedFV, formattedRFV)

	// Multi-line formatted printing
	fmt.Printf(`
	Future Value: %.2f
	Future Value (adjusted for inflation): %.2f`, futureValue, futureRealValue)

}

func printText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)	
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return
}