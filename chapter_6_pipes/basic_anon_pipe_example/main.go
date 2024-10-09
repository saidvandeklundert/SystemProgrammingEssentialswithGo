package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Create and run the echo command
	echoCmd := exec.Command("echo", "Hello, world!")
	echoOutput, err := echoCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running echoCmd: %v\n", err)
		return
	}

	// Create the grep command with the output of echoCmd as its input
	grepCmd := exec.Command("grep", "Hello")
	grepCmd.Stdin = strings.NewReader(string(echoOutput))

	// Capture the output of grepCmd
	grepOutput, err := grepCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running grepCmd: %v\n", err)
		return
	}

	// Print the output of grepCmd
	fmt.Printf("Output of grep: %s", grepOutput)
}

/*
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	echoCmd := exec.Command("echo", "Hello, world!")
	grepCmd := exec.Command("grep", "Hello")

	pipe, err := echoCmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating StdoutPipe for echoCmd: %v\n", err)
		return
	}

	if err := grepCmd.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting grepCmd: %v\n", err)
		return
	}

	if err := echoCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running echoCmd: %v\n", err)
		return
	}

	if err := pipe.Close(); err != nil {
		fmt.Fprintf(os.Stderr, "Error closing pipe: %v\n", err)
		return
	}

	if err := grepCmd.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "Error waiting for grepCmd: %v\n", err)
		return
	}
}
*/
