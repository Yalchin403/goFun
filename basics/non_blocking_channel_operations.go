package main

import (
	"fmt"
	"time"
)

func handler(c chan string, message string) {
	time.Sleep(2 * time.Second)
	c <- message
}

func main() {
	channel := make(chan string)
	go handler(channel, "Message for the channel")

	select {
	case msg := <-channel:
		fmt.Println(msg)

	default:
		fmt.Println("No message received, timeout!")
	}
}
