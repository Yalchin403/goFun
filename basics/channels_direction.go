package main

/*
When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
This specificity increases the type-safety of the program.
*/

func ping(ping chan<- string, message string) {
	ping <- message
	// here ping receives the message, so it's receiver not sender
}

func pong(pong chan<- string, ping <-chan string) {
	// here pong is the receiver, ping is the sender
	msg := <-ping
	pong <- msg
}
