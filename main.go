package main

import (
	"errors" // We need this built-in package to create custom errors
	"fmt"
)

// 1. MULTIPLE RETURN VALUES
// This function promises to return TWO things: a string message, and an error.
func pingServer(ip string) (string, error) {

	// We check for bad input first.
	if ip == "" {
		// We return a blank string, and a newly created error.
		return "", errors.New("CRITICAL: IP address cannot be completely blank")
	}

	if ip == "0.0.0.0" {
		return "", errors.New("WARNING: 0.0.0.0 is an invalid routing address")
	}
	if ip == "127.0.0.1" {
		return "", errors.New("ERROR: Cannot ping the local loopback address.")
	}

	// If the CPU makes it past the checks, the IP is valid.
	// We return our success string, and 'nil' (meaning absolutely no error).
	return "Ping successful! Packet returned from " + ip, nil
}

func main() {
	fmt.Println("--- NETWORK DIAGNOSTIC TOOL ---")

	// 2. THE TARGET
	targetIP := "127.0.0.1"

	// 3. CAPTURING MULTIPLE RETURNS
	// We catch the string in 'result', and the error in 'err'
	result, err := pingServer(targetIP)

	// 4. THE GO ERROR CHECK
	// "If the error is NOT equal to nothing..." (Meaning, an error exists!)
	if err != nil {
		// Print the error and immediately stop the program using 'return'
		fmt.Println("DIAGNOSTIC FAILED:", err)
		return
	}

	// 5. SUCCESS PATH
	// If the CPU reaches here, we know 100% that err was nil.
	fmt.Println("DIAGNOSTIC PASSED:", result)
}
