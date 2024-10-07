package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	// create a buffered channel for OS signals
	signals := make(chan os.Signal, 1)
	// create a channel that can be used to indicate the program is done
	done := make(chan struct{}, 1)

	// registers the os.Interrupt signal with the signals channel
	signal.Notify(signals, os.Interrupt)
	// a gorouting is started to run concurrently with the main program
	go func() {
		for {
			s := <-signals
			switch s {
				// when the signal is 'os.Interrupt', send an empty struct
				// to indicate we are done
			case os.Interrupt:
				fmt.Println("INTERRUPT")
				done <- struct{}{}
			default:
				fmt.Println("OTHER")
			}

		}
	}()

	fmt.Println("awaiting signal")
	// <-done blocks until a value is received from the done channel,
	<-done
	fmt.Println("exiting")
}
