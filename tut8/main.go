package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

// Working with print formatting
func main() {
	fmt.Printf("%s %d %c %f %t %o %x\n",
		"Stuff", 1, 'A', 3.14, true, 1, 1)

	// Formatting of floating point numbers
	fmt.Printf("%9f\n", 3.14)
	fmt.Printf("%.2f\n", 3.141592)
	fmt.Printf("%.1f\n", 3.141592)
	fmt.Printf("%9.f\n", 3.141592)
	fmt.Printf("%.f\n", 3.141592)

	sp1 := fmt.Sprintf("%9.f", 3.141592)
	sp2 := fmt.Sprintf("%c", 'A')
	pl("Type of data: ", reflect.TypeOf(sp1))
	pl("Type of data: ", reflect.TypeOf(sp2))
}
