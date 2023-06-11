package main

import (
	"fmt"
)

var pl = fmt.Println

func getAverage(nums ...float64) float64 {

	sum := 0.0
	var numSize = float64(len(nums))

	for _, val := range nums {
		sum += val
	}

	return sum / numSize

}

func main() {

	slNum := []float64{11, 13, 17}
	pl("running: ")
	fmt.Printf("Average: %.3f", getAverage(slNum...))

}
