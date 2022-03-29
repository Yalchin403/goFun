package main

import (
	"fmt"
	"strconv"
	"sync"
)

//func asyncFunction(randomVal int, c chan int) {
//
//	// send it to the channel
//	c <- randomVal * 12
//}
//
//func main() {
//	channel := make(chan int)
//	go asyncFunction(12, channel)
//	// receive from the channel
//
//	messageFromFunc := <-channel
//	fmt.Println(messageFromFunc)
//}
var waitGroup sync.WaitGroup

func foo(val string, c chan string) {
	defer waitGroup.Done()
	c <- val
}

func main() {
	channel := make(chan string, 10)

	for i := 0; i < 10; i++ {
		conc := "Coder403" + strconv.Itoa(i)
		waitGroup.Add(1)
		go foo(conc, channel)
	}
	waitGroup.Wait()
	close(channel)

	for item := range channel {
		fmt.Println(item)
	}

}
