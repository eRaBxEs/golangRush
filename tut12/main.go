package main

import (
	"fmt"
)

// working with slice
var pl = fmt.Println

func main() {
	// var name []dataType
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "simulated"
	sl1[4] = "golang"
	sl1[5] = "universe"

	pl("Length of the slice:", len(sl1))

	// loop 1:
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}

	// loop 2:
	for _, x := range sl1 {
		pl(x)
	}

	// creating slices from an array
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl("first 3:", sArr[0:3])
	pl("last 3:", sArr[2:])
	pl("sl3 before: ", sl3)
	sArr[0] = 10
	pl("sl3 after :", sl3)

	// appending to a slice created from an array
	sl3 = append(sl3, 12)
	pl("sl3 after appending: ", sl3)
	pl("sArr after appending to sl3: ", sArr)

}
