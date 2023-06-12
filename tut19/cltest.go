package main

import (
	"fmt"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	pl("look carefully!")
	fmt.Println(os.Args)
	args := os.Args[1:]
	var iArgs []int
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}
	max := 0

	for _, x := range iArgs {
		if x > max {
			max = x
		}
	}
	pl("Max Value:", max)
}
