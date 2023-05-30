package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	// In Go, characters are called runes
	// And runes are unicodes that represents character
	rStr := "abcdefg"
	pl("Rune Count:", utf8.RuneCountInString(rStr))
	// loop through each rune
	for i, runeVal := range rStr {
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
