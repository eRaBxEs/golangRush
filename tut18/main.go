package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	// check for error if file exists or not
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't exist!")
	} else {
		f, err := os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}

}
