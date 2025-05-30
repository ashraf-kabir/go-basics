package main

import (
	"fmt"
	"price-calulcator/filemanager"
	"price-calulcator/prices"
)

func main() {
	taxRates := []float64{0.4, 0.5, 0.6}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		go priceJob.Process(doneChans[i], errorChans[i])

		// if err != nil {
		// 	fmt.Println("Could not process price job")
		// 	fmt.Println(err)
		// }
	}

	for i, _ := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				panic(err)
			}
		case <-doneChans[i]:
			fmt.Printf("done %v\n", i)
		}
	}
}
