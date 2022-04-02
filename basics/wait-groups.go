package main

import (
	"fmt"
	"strconv"
	"sync"
)

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

	results := make([]string, 3)

	for item := range channel {
		results = append(results, item)
	}

	fmt.Println(results)
}
