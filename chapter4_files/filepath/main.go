package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir := "home/klundert"
	file := "document.txt"

	fullPath := filepath.Join(dir, file)
	fmt.Println("Full path:", fullPath)
}
