package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubledNumbers := transformNumbers(&numbers, double)
	fmt.Println(doubledNumbers)
	tripledNumbers := transformNumbers(&numbers, triple)
	fmt.Println(tripledNumbers)
}

// function type on transform argument
func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumbers []int
	for _, v := range *numbers {
		dNumbers = append(dNumbers, transform(v))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
