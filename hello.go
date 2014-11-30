package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Set the who in this case
	who := "World!"

	// If more than 1 command-line argument is passed...
	if len(os.Args) > 1 {
		// Join together the command-line arguments with spaces, save to the who variable
		who = strings.Join(os.Args[1:], " ")
	}

	fmt.Println("Hello", who)
}