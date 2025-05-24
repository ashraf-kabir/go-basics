package main

import "fmt"

func main() {
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sum := sumup(numbers)

	sum := sumup(1, 2, 3)

	fmt.Println(sum)
}

// variadic function that takes any number of args
func sumup(numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}
