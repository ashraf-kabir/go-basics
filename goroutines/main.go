package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Greet:", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Greet:", phrase)
}

func main() {
	go greet("Hello World")
	go greet("How are you?")
	go slowGreet("How... are... you...?")
	go greet("I hope you are doing well.")
}
