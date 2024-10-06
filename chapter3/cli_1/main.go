package main

/*
	go run main.go word1 word2 word3 > stdout.txt 2> stderr.txt
*/
import (
	"fmt"

	"os"
)

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		fmt.Fprintln(os.Stderr, "No words provided.")
		os.Exit(1) // 0 is success, 1 is failure in exit-code-speak
	}

	app(words)

}
func app(words []string) {
	for _, w := range words {
		if len(w)%2 == 0 {
			fmt.Fprintf(os.Stdout, "word %s is even\n", w)
		} else {
			fmt.Fprintf(os.Stderr, "word %s is odd\n", w)
		}
	}
}
