package main

import (
	"fmt"
	"strings"
)

var pl = fmt.Println

func main() {
	// A string is simply an array of bytes i.e []byte
	sV1 := "A word"

	// creating a replacer
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl("Length of sV2", len(sV2))
	pl("sV2 Contains Another:", strings.Contains(sV2, "Another"))
	pl("o index in sV2:", strings.Index(sV2, "o"))
	pl("replace o with 0", strings.Replace(sV2, "o", "0", -1))
	sV3 := "\nSome Words\n"
	pl(sV3)
	sV3 = strings.TrimSpace(sV3)
	pl(sV3)
	pl("split :", strings.Split("a-b-c-d", "-"))
	pl("To Lower:", strings.ToLower(sV2))
	pl("To Upper:", strings.ToUpper(sV2))
	pl("Prefix:", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix:", strings.HasSuffix("tacocat", "cat"))

}
