package main

import (
	"fmt"
	"os"
)

func main() {
	info, err := os.Stat("main.gso")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		} else {
			panic(err)
		}

	}
	fmt.Printf("File name: %s\n", info.Name())
	fmt.Printf("File size: %d bytes\n", info.Size())
	fmt.Printf("File permissions: %s\n", info.Mode())
	fmt.Printf("Last modified: %s\n", info.ModTime())
}
