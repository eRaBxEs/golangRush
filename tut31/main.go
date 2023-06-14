package main

import (
	"fmt"
)

var pl = fmt.Println

// using a function to call a pass a parameter of type func
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", f(x, y))
}

func sumVal(x, y int) int {
	return x + y
}

func main() {
	// closure are functions that don't have to be associated with an identifier,
	// but very often they are associated with a variable.
	intSum := func(x, y int) int { return x + y }
	pl("5 + 4 =", intSum(5, 4))

	// closure can change values outside of a function
	samp1 := 1
	changeVar := func() {
		samp1 += 1
	}

	changeVar() // calling the closure

	pl("samp1 = ", samp1)

	useFunc(sumVal, 5, 8)
}
