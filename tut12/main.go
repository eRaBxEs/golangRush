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
}
