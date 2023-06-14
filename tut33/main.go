package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	reStr := "The ape was at the apex"
	match, _ := regexp.Match("(ape[^ ]?", []byte(reStr)) // look for ape without space after ape
	pl(match)
}
