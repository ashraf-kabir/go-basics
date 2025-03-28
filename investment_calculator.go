package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	// fmt.Print("Invested amount: ")
	outputText("Invested amount: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedStr := fmt.Sprintf("Future value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Println("Future value: ", futureValue)\
	// fmt.Printf("Future value: %v\nFuture value (adjusted for inflation): %v\n", futureValue, futureRealValue)
	fmt.Print(formattedStr)
	// fmt.Println("Future value (adjusted for inflation): ", futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}