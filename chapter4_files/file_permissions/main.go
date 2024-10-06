package main

import (
	"fmt"
	"os"
)

func main() {
	// Stat the file to get its information
	fileInfo, err := os.Stat("main.go")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// get the file permissions:
	permissions := fileInfo.Mode().Perm()
	permissionString := fmt.Sprintf("%o", permissions)
	fmt.Printf("Permissions: %s\n", permissionString)
}
