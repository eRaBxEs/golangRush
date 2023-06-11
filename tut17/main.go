package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

// Working with files
func main() {
	// create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	iPrimeArr := []int{2, 3, 5, 7, 11}
	var sPrimeArr []string
	for _, i := range iPrimeArr {

		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	// write to the file
	for _, st := range sPrimeArr {
		_, err := f.WriteString(st + "\n")
		if err != nil {
			log.Fatal((err))
		}
	}

	// open the file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() { // read the file on a line by line basis
		pl("Prime: ", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
