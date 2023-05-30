package main

import (
	"fmt"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	// Generate a random seconds integer using time that elapsed from Jan 1, 1970
	seedSecs := time.Now().Unix()
	// Then seed it using the random feature from the math/rand package
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random number:", randNum)

}
