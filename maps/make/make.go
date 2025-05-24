package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//userNames := []string{}
	userNames := make([]string, 2, 5) // type, length, capacity

	userNames[0] = "Jason"

	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")

	fmt.Println(userNames)

	//courseRatings := map[string]float64{}
	//courseRatings := make(map[string]float64, 3)
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.5
	courseRatings["java"] = 4.8
	courseRatings["php"] = 5.00

	//fmt.Println(courseRatings)
	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
