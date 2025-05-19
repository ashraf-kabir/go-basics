package main

import "fmt"

func main() {
	var productNames [3]string
	productNames[0] = "A"
	productNames[1] = "B"
	productNames[2] = "C"

	prices := [3]float64{1.99, 2.99, 3.99}

	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(len(prices))
	fmt.Println(productNames)
	fmt.Println(len(productNames))

	// Slices
	featuredPrices := prices[0:2] // take index 0 & 1
	fmt.Println(featuredPrices)
}
