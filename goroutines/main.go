package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Greet:", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Greet:", phrase)
	doneChan <- true
	close(doneChan)
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("Hello World", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How... are... you...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you are doing well.", done)

	// for _, d := range dones {
	// 	<-d
	// }

	// <-done

	for range done {
		// fmt.Println(doneChan)
	}
}
