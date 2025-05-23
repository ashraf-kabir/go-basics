package main

import "fmt"

func main() {
	var productNames [3]string
	productNames[0] = "A"
	productNames[1] = "B"
	productNames[2] = "C"

	prices := [3]float64{1.99, 2.99, 3.99}

	prices[1] = 4.99 // changing the value of prices[1] will also change the value of featuredPrices2 and featuredPrices3

	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(len(prices))
	fmt.Println(productNames)
	fmt.Println(len(productNames))

	// Slices
	featuredPrices1 := prices[0:2] // start from index 0 and ends before index 2
	featuredPrices2 := prices[:2]  // start from beginning and ends before index 2
	featuredPrices3 := prices[1:]  // start index 1 and ends at the end

	fmt.Println(featuredPrices1)
	fmt.Println(featuredPrices2)
	fmt.Println(featuredPrices3)

	fmt.Println(len(prices)) // length of the array
	fmt.Println(cap(prices)) // capacity of the array

	featuredPrices1 = featuredPrices1[:3]

	fmt.Println(len(featuredPrices1)) // length of the slice
	fmt.Println(cap(featuredPrices1)) // capacity of the slice
}
