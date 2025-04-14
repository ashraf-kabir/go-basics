package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // address of age, pointer

	fmt.Println("age: ", age)
	fmt.Println("agePointer: ", agePointer) // prints age pointer
	fmt.Println("age: ", *agePointer) // prints age, de-reference

	adultYears := getAdultYears(&age)
	fmt.Println(adultYears)
}

// avoiding a copy of integer by using pointer
func getAdultYears(age *int) int {
	return *age - 18
}