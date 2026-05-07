package main

import (
	"fmt"
	"os"
)

func main() {
	// len() checks how many items are in os.Args
	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a username.")
		fmt.Println("Usage: go run main.go [name]")
		os.Exit(1) // Exit with 1 to signal a failure to the OS
	}

	fmt.Println("System Online: Hello", os.Args[1])
}
