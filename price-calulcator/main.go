package main

import (
	"fmt"
	"price-calulcator/filemanager"
	"price-calulcator/prices"
)

func main() {
	taxRates := []float64{0.4, 0.5, 0.6}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}
}
