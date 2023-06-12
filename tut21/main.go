package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// working with map
	// myMap map[keyType] valueType

	// way 1:
	var heroes map[string]string
	heroes = make(map[string]string)

	// way 2:
	villains := make(map[string]string)

	// usage
	heroes["Batman"] = "Bruce Wayne"
	heroes["Suoerman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villains["Lex Luther"] = "Lex Luther"

	superPets := map[int]string{
		1: "Krypto",
		2: "Bat Hound",
	}

	fmt.Printf("Batman is %v\n", heroes["Batman"])
	pl("Chip :", superPets[3])
	// checking for the existence of a 3rd pet
	_, ok := superPets[3]
	pl("Is there a 3rd pet: ", ok)

	// loop thru a map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// delete a map
	delete(heroes, "The Flash")

}
