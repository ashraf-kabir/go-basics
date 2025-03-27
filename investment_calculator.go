package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Invested amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedStr := fmt.Sprintf("Future value: %.2f\nFuture value (adjusted for inflation): %.2f\n", futureValue, futureRealValue)
	// fmt.Println("Future value: ", futureValue)\
	// fmt.Printf("Future value: %v\nFuture value (adjusted for inflation): %v\n", futureValue, futureRealValue)
	fmt.Print(formattedStr)
	// fmt.Println("Future value (adjusted for inflation): ", futureRealValue)
}
