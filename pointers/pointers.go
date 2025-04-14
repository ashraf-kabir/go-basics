package main

import "fmt"

func main() {
	age := 32 // regular variable

	agePointer := &age // address of age, pointer

	fmt.Println("age: ", age)
	fmt.Println("agePointer: ", agePointer) // prints age pointer
	fmt.Println("age: ", *agePointer) // prints age, de-reference

	// adultYears := getAdultYears(&age)
	// fmt.Println(adultYears)
	getAdultYears(&age)
	fmt.Println("age: ", age)
}

// avoiding a copy of integer by using pointer
func getAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}