package main

import "fmt"

func main() {
	prices := []float64{10, 20, 30}
	taxRates := []float64{0.4, 0.5, 0.6}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		var taxIncludedPrices []float64 = make([]float64, len(prices))
		for j, price := range prices {
			taxIncludedPrices[j] = price * (1 + taxRate)
		}
		result[taxRate] = taxIncludedPrices
	}

	fmt.Println(result)
}
