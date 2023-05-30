package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	// converting a string numeral to int
	cV1 := "70000"
	cv2, err := strconv.Atoi(cV1)
	if err == nil {
		pl(cv2, err, reflect.TypeOf(cv2))
	}

	// converting an int to a string
	cv3 := 8000
	cv4 := strconv.Itoa(cv3)
	if err == nil {
		pl(cv4)
	}

	// converting a string numeral to float
	cv5 := "34.56"
	if cv6, err := strconv.ParseFloat(cv5, 64); err == nil {
		pl(cv6, reflect.TypeOf(cv6))
	}

	// converting a float into a string
	cv7 := 3.14
	cv8 := fmt.Sprintf("%f", cv7)
	pl(cv7, reflect.TypeOf(cv8))
}
