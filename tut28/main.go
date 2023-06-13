package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2", i)
	}
}

func main() {

	// using go routines
	go printTo10()
	go printTo15()

	// In order to slow down the main func, we use a sleep timer
	time.Sleep(2 * time.Second)
}
