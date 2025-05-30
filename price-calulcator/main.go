package main

import (
	"fmt"
	"price-calulcator/filemanager"
	"price-calulcator/prices"
)

func main() {
	taxRates := []float64{0.4, 0.5, 0.6}
	doneChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process price job")
		// 	fmt.Println(err)
		// }
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
