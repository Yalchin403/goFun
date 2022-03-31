package main

import "fmt"

/*
Closing a channel indicates that no more values will be
sent on it. This can be useful to communicate completion to
the channel’s receivers.
*/

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs // returns value of j, and true if any data received
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs) // sends 0, false to the channel with which we can stop 1st for loop
	fmt.Println("sent all jobs")

	fmt.Println(<-done)
	close(done)
}
