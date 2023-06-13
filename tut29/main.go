package main

import (
	"fmt"
)

var pl = fmt.Println

func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

func main() {
	// this will enable us to get the values sent to the channels in order
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	// print them in order
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)

}
