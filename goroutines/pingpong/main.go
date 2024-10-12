package main

import (
	"fmt"
	"strings"
)

func shout(ping chan string, pong chan string) {
	for {
		message := <-ping
		pong <- strings.ToUpper(message)
	}

}

func main() {
	fmt.Println("Explaining channels through a super simple script.")
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	fmt.Println("Shout something, press 'q' to quit.")
	for {
		fmt.Printf(">")
		var user_input string
		fmt.Scanln(&user_input)
		if user_input == "q" {
			break
		}
		ping <- user_input
		result := <-pong
		fmt.Println(result)
	}
	close(pong)
	close(ping)
	fmt.Println("bye")
}
