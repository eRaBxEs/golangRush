package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	aStr1 := "abcd"
	// quickly convert a string into a slice of rune
	rArr := []rune(aStr1)
	// now to print each element of the rune array/slice
	for _, v := range rArr {
		fmt.Printf("Rune array : %d\n", v)
	}

	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:]) // byteArr[:] creates a fresh slice and not a representation that references byteArr
	pl("I am a string:", bStr)

}
