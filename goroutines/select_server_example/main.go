package main

import (
	"fmt"
	"time"
)

func serverOne(channel chan string) {
	for {
		time.Sleep(3 * time.Second)
		channel <- "message from ServerOne"
	}

}

func serverTwo(channel chan string) {
	for {
		time.Sleep(1 * time.Second)
		channel <- "message from ServerTwo"
	}

}

func main() {
	serverChannelOne := make(chan string)
	serverChannelTwo := make(chan string)

	go serverOne(serverChannelOne)
	go serverTwo(serverChannelTwo)

	for {
		select {
		case one := <-serverChannelOne:
			fmt.Println(one)

		case two := <-serverChannelTwo:
			fmt.Println(two)

		}
	}

}
