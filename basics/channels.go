package main

import "fmt"

func asyncFunction(randomVal int, c chan int) {

	// send it to the channel
	c <- randomVal * 12
}

func main() {
	channel := make(chan int)
	go asyncFunction(12, channel)
	// receive from the channel

	messageFromFunc := <-channel
	fmt.Println(messageFromFunc)
}
