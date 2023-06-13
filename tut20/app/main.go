package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"log"
	"reflect"
)

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Print(strArr)
	fmt.Println(reflect.TypeOf(strArr))

	// demonstrating what encapsulation is in Golang
	date := stuff.Date{}
	err := date.SetDay(21)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1974)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("1st Day : %d/%d/%d", date.Day(), date.Month(), date.Year())

}
