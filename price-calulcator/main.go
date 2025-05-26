package main

import (
	"price-calulcator/prices"
)

func main() {
	taxRates := []float64{0.4, 0.5, 0.6}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}
}
