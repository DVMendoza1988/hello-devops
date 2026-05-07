package main

import (
	"fmt"
	"os"
	"strconv" // New Package: String Conversion
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [number_of_servers]")
		os.Exit(1)
	}

	// 1. Arguments always come in as Strings.
	//2. We must convert the string to an Integer (int) to do math.
	input := os.Args[1]
	serverCount, err := strconv.Atoi(input)

	// In go, we handle errors IMMEDIATELY.

	if err != nil {
		fmt.Println("Error: Please provide a whole number.")
		os.Exit(1)
	}

	const coresPerServer = 4 // Constants are declared with 'const' and cannot be changed.
	totalCores := serverCount * coresPerServer
	fmt.Printf("Infratructure report:\n")
	fmt.Printf("- Servers %d\n", serverCount)
	fmt.Printf("- Total CPU Cores Needed: %d\n", totalCores)
}
