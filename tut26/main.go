package main

import (
	"fmt"
)

var pl = fmt.Println

// Defined types in Golang

type TsP float64
type TBs float64
type ML float64

func TspToML(tsp TsP) ML {
	return ML(tsp * 4.92)
}

func TBsToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Using Associate methods
func (tsp TsP) ToML() ML {
	return ML(tsp * 4.92)
}

func (tbs TBs) ToML() ML {
	return ML(tbs * 14.79)
}

func main() {

	ml1 := ML(TsP(3) * 4.92)
	fmt.Printf("3 tsp = %.2f ML\n", ml1)
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 tbs = %.2f ML\n", ml2)

	pl("2 tsp + 4 tbs = ", TBs(2), TBs(4))
	pl("2 tsp > 4 tsp = ", TBs(2) > TBs(4))

	fmt.Printf("3 tsp = %.2f ML\n", TspToML(3))
	fmt.Printf("3 tbs = %.2f ML\n", TBsToML(3))

	tsp1 := TsP(3)
	tbs1 := TBs(2)
	// utilising the associated method
	fmt.Printf("%.2f tsp = %.2f ML\n", tsp1, tsp1.ToML())
	fmt.Printf("%.2f tbs = %.2f ML\n", tbs1, tbs1.ToML())

}
