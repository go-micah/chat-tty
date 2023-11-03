package main

import (
	"bufio"
	"fmt"
	"os"
)

func stringPrompt(label string) string {

	var s string
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}

	return s
}

func main() {

	fmt.Printf("Hi there. You can ask me stuff!\n")

	for {
		fmt.Printf("You said: %v", stringPrompt(">"))
	}
}
