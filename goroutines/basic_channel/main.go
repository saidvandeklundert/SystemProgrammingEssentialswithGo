package main

import "fmt"

func main() {
	// create a channel of the type string:
	messages := make(chan string)

	// send in a string:
	go func() { messages <- "ping" }()

	// read from the channel:
	msg := <-messages
	fmt.Println(msg)
	// send in another string:
	go func() { messages <- "pong" }()
	// read from the channel agaon:
	msg = <-messages
	fmt.Println(msg)
	// close the channel
	close(messages)

	// create a buffered channel with a capacity of 2:
	buffered_messages := make(chan string, 2)

	buffered_messages <- "buffered ping"
	buffered_messages <- "buffered pong"

	fmt.Println(<-buffered_messages)
	fmt.Println(<-buffered_messages)

}
