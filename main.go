package main

import (
	"fmt"
	"time"
)

// 1. THE CONCURRENT WORKER
// We add a new paramater: 'results'
// 'chain string' stricktly means a "a Channel that only accepts strings."
func checkServerHealth(serverName string, results chan string) {
	time.Sleep(1 * time.Second) // Simulate network work

	// 2. SENDING DATA INTO A CHANNEL
	// The arrow <- points INTO the channel. We are shoving this text down the pipe.
	results <- "[OK] " + serverName + " is online."
}

func main() {
	servers := []string{"Delta-Auth", "Alpha-DB", "Beta-Cache", "Gamma-Web"}

	// 3. CREATING THE CHANNEL
	// We must use 'make()' command to physically allocate the channel in memory.
	healthStream := make(chan string)

	fmt.Println("--- DISPATCHING PROBES ---")

	// 4. SPAWNING GOROUTINES
	for _, name := range servers {
		// We pass our channel into each background worker.
		// They all share the exact same pipe!
		go checkServerHealth(name, healthStream)
	}

	// 5. RECEIVING FROM THE CHANNEL
	// We know we sent exactly 3 probes (len(servers)),
	// so we need to pull exactly 3 results off the conveyor belt.
	for i := 0; i < len(servers); i++ {
		// The arrow <- is IN FRONT of the channel. we are pulling data OUT of the channel.
		// The CPU will pause on this line and wait if the pipe is currently empty.
		msg := <-healthStream
		fmt.Println("Received: ", msg)
	}

	fmt.Println("--- ALL DATA AGGREGATED ---")
}
