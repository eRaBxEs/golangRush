package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	restr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")
	pl("MatchString :", r.MatchString(restr2))
	pl("FindString :", r.FindString(restr2))
	pl("Index :", r.FindStringIndex(restr2))
	pl("All String :", r.FindAllString(restr2, -1))
	pl("1st two Strings :", r.FindAllString(restr2, 2))
	pl("All Submatch Index :", r.FindAllStringSubmatch(restr2, -1))
	pl(r.ReplaceAllString(restr2, "Dog"))

}
