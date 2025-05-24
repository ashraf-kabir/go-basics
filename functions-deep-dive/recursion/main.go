package main

import "fmt"

func main() {
	fmt.Println(factorial(5))
}

func factorial(number int) int {
	//result := 1
	//for i := 1; i <= number; i++ {
	//	result = result * i
	//}
	//return result

	// recursive approach
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

// factorial of 5: 5 * 4 * 3 * 2 * 1 = 120
