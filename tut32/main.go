package main

import (
	"fmt"
)

var pl = fmt.Println

// Recursion

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	pl("Factorial of 4 = ", factorial(4))
}
