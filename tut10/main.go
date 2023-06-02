package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// working with arrays in Golang
	var arr1 [5]int // create an empty array using var
	arr1[0] = 1     // assign a value to the first index in an array
	// creating an array with values
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0 of arr2:", arr2[0])
	pl("Length of arr2:", len(arr2))
	// looping through an array using the traditional for loop construct
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	// looping through an array using range for loop construct
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// working with muti-dimensional arrays
	// Sample 1: 2 by 2 dimensional array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	// looping through the elements of a 2 by 2 dimensional array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	arr4 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	// looping through the elements of a 3 by 3 dimensional array
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			pl(arr4[i][j])
		}
	}

}
