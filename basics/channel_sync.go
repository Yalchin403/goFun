package main

import (
	"fmt"
	"time"
)

func worker(c chan bool) {
	time.Sleep(2)
	c <- true
}

func main() {
	channel := make(chan bool)
	go worker(channel)

	// wait for notification from worker
	notification := <-channel
	fmt.Println(notification)
}
