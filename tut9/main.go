package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Guessing game

var pl = fmt.Println

func main() {

	// Get a seed
	seedSec := time.Now().Unix()
	// Sow the seed
	rand.Seed(seedSec)
	// Get a random number from 1 - 50 inclusive
	randNum := rand.Intn(50) + 1
	pl("The random number is ", randNum)
	for true {
		pl("Guess a number between 1 and 50 inclusive:")
		// create a reader from the bufio package
		reader := bufio.NewReader(os.Stdin)
		// call the ReadString method of the reader you created earlier
		// use a delimiter to determine the extent to which the string from the console should be read
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if iGuess < randNum {
			pl("Pick a higher value")
		} else if iGuess > randNum {
			pl("Pick a lower value")
		} else {
			pl("You guessed right")
			break
		}
	}

}
