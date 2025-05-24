package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// sum := sumup(numbers)

	sum := sumup(1, 2, 3)
	anotherSum := sumup(1, numbers...)

	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// variadic function that takes any number of args
// first arg added as example & not used
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, v := range numbers {
		sum += v
	}

	return sum
}
