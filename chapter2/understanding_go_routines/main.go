package main

import (
	"fmt"
	"sync"
)

// print a string 5 times and calls Done on the wg
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 5; i++ {
		fmt.Println(s)
	}
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	fmt.Println("Hello world!")
	// without the wait groups, the main thread would close without waiting for any of the created
	// Go-routines to finish.
	go say("hello", &wg)
	go say("world", &wg)
	wg.Wait()
}
