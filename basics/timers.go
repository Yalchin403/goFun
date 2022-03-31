package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C // returns time when it is fired
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	var input int
	fmt.Println("Please enter a number: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}

	var stop2 bool
	if input == 5 {
		stop2 = timer2.Stop() // it will not fire if the input is 5, we prove that by sleeping for 2 sec at the end
		// line 18 will never be executed if the input is 5
	}

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
