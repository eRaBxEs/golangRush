package main

import (
	"fmt"
)

var pl = fmt.Println

type contact struct {
	fName string
	lName string
	phone string
}

// using structs for composition
type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at company %s is %s %s", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	con1 := contact{
		"James",
		"Wang",
		"555-2121",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North st",
		con1,
	}

	bus1.info()

}
