package main

import (
	"fmt"
)

var pl = fmt.Println

type Animal interface {
	AngrySound()
	HappySound()
}

type Cat string // creating a type Cat

func (c Cat) Attack() { // associating the type to a function
	pl("Cat attacks it's prey")
}

func (c Cat) Name() string {
	return string(c)
}

// Now to make the cat implement the interface
func (c Cat) AngrySound() {
	pl("Cat says Hissssss!")
}
func (c Cat) HappySound() {
	pl("Cat says Purrrrr!")
}

func main() {
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	var kitty2 = kitty.(Cat) // interface type assertion
	kitty2.Attack()
	pl("Cat's Name is :", kitty2.Name())
}
