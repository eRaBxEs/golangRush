package main

import (
	"fmt"
)

var pl = fmt.Println

// Working with variadic function
func getSum(nums ...int) int {
	sum := 0

	for _, x := range nums {
		sum += x
	}

	return sum
}

// function that accepts a slice as an argument
func getSumArray(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {

	pl("Sum of numbers:", getSum(1, 2, 3, 4, 5))

	vArr := []int{1, 2, 3, 4, 5}

	pl("Sum of an array:", getSumArray(vArr)) // here the slice is passed by value because the function does not change the slice

}
