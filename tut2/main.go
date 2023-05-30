package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	// variables in Go are immutable,
	// meaning that they can't change to contain values of another data type
	var vName = "ehioje"
	var v1, v2 = 1.2, 3.4
	var v3 = "hello"
	v4 := 3.4
	pl(vName, v1, v2, v3, v4)

	// showing the data types of various literals
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('\n'))
}
