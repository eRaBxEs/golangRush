package main

import (
	"fmt"
)

var pl = fmt.Println

// here the int is passed by value because the function does not change the int
func changeVal(f3 int) int {
	f3 += 7
	return f3
}

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {

	f3 := 5
	pl("f3 before func:", f3)
	_ = changeVal(f3)
	pl("f3 after func:", f3)

	f4 := 10
	pl("f4 before func:", f4)
	changeVal2(&f4)
	pl("f4 after func:", f4)

	f5 := 8
	var f5Ptr *int = &f5
	pl("f5 Address:", f5Ptr)
	pl("f5 Value:", *f5Ptr)
	*f5Ptr = 11
	pl("f5 Address after :", f5Ptr)
	pl("f5 Value after :", *f5Ptr)

	pArr := [4]int{1, 2, 3, 4}
	pl("pArr before func: ", pArr)
	dblArrVals(&pArr)
	pl("pArr after func: ", pArr)

}
