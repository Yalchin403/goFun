package main

import (
	"fmt"
	"time"
)

/*
Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select
is a powerful feature of Go.
*/

func goToSleep(c chan int, val int) {
	time.Sleep(time.Duration(val) * time.Second)
	c <- val
}

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go goToSleep(channel1, 1)
	go goToSleep(channel2, 2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Message received from the first channel:", msg1)

		case msg2 := <-channel2:
			fmt.Println("Message received from the second channel:", msg2)

		}
	}
}
