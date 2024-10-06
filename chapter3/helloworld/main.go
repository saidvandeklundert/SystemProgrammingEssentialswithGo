package main

import (
	"fmt"
	"unsafe"

	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func main() {
	fmt.Println("Hello World!")
	unix.Syscall(unix.SYS_WRITE, 1, uintptr(unsafe.Pointer(&[]byte("Hello, World!")[0])),
		uintptr(len("Hello, World!")),
	)
	// start another process
	cmd := exec.Command("ls", "-ltr")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the current process ID
	pid := os.Getpid()
	fmt.Println("Current process ID:", pid)

}
