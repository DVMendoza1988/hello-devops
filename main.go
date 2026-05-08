package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. The Safety Gate (Prevents a crash if you forget to type a word)
	if len(os.Args) < 2 {
		fmt.Println("ERROR: Please provide a service name.")
		fmt.Println("Usage: go run main.go [service]")
		os.Exit(1)
	}

	// 2. Grab the word you typed in the terminal
	targetService := os.Args[1]

	// 3. Our internal database
	portDB := map[string]int{
		"ssh":     22,
		"http":    80,
		"dns":     53,
		"https":   443,
		"ftp":     21,
		"smtp":    25,
		"pop3":    110,
		"imap":    143,
		"ldap":    389,
		"rdp":     3389,
		"sql":     1433,
		"redis":   6379,
		"mongodb": 27017,
		"kafka":   9092,
		"telnet":  23,
		"ntp":     123,
		"IMAP":    143,
		"imaps":   993,
	}

	// 4. Look up whatever word you typed into the terminal
	portNumber, exists := portDB[targetService]

	// 5. The Output
	if exists == true {
		fmt.Printf("[RESOLVED] Service '%s' operates on Port %d\n", targetService, portNumber)
	} else {
		fmt.Printf("[ERROR] Unknown service: '%s'\n", targetService)
	}
}
