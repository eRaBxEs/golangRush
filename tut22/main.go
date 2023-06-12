package main

import (
	"fmt"
)

var pl = fmt.Println

// The generic type paramter is going to be capitalized
// and it also has to be between square brackets

type MyConstraint interface {
	int | float64
}

// pkg.go.dev/golang.org/x/exp/constraints  : for more on constraints

// Using Generic type
func getSum[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5 + 4 = ", getSum(5, 4))
	pl("5.6 + 4.7 = ", getSum(5.6, 4.7))

}
