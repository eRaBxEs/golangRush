package main

import (
	"fmt"
)

var pl = fmt.Println

type rectangle struct {
	length, breadth float64
}

func (r rectangle) Area() float64 {
	return r.length * r.breadth
}

func main() {

	rect1 := rectangle{5.0, 6.2}
	pl("Rect area: ", rect1.Area())

}
