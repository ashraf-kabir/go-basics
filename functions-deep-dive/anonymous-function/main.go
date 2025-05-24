package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// anonymous function called on the 2nd arg
	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	fmt.Println(transformed)
}

// function type on the 2nd arg
func transformNumbers(numbers *[]int, transform func(int) int) []int {
	var dNumbers []int
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}
