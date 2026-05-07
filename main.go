package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing argument.")
		os.Exit(1)
	}

	// Declare a variable named 'user' and assign it the value from Args
	var user string = os.Args[1]

	fmt.Printf("Access Granted: Welcome, %s\n", user)
}
