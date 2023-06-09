package main

import (
	"fmt"
)

var pl = fmt.Print

// looking at struct

type customer struct {
	name    string
	address string
	bal     float64
}

func getCustInfo(c customer) {
	fmt.Printf("%s owes us %f\n", c.name, c.bal)
}

func newCustAddress(c *customer, address string) {
	c.address = address
}

func main() {
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "9 Alasco avenue"
	tS.bal = 234.56

	getCustInfo(tS)

	newCustAddress(&tS, "77 heaven st")
	// check if address has truely changed
	pl("Address: ", tS.address, "\n")

	sS := customer{"Sally Smith", "123 Main st", 0.0}
	pl("Name: ", sS.name)

}
