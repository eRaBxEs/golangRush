package main

import (
	"fmt"
)

var pl = fmt.Println

func getSum(x int, y int) int {
	return x + y
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("you can't divide by zero")
	} else {
		return x / y, nil
	}
}

func main() {
	pl("getSum :", getSum(3, 4))
	a, b := getTwo(3)
	fmt.Printf("getTwo: %d, %d", a, b)

	pl(getQuotient(5, 0))
	pl(getQuotient(5, 4))
}
