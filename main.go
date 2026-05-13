package main

import (
	"fmt"
	"sync"
	"time" // Gives the ability to pause the computer
)

// 1. THE CONCURRENT WORKER
// We pass in our server name, and a Pointer to our WaitGroup counter.
func checkServerHealth(serverName string, wg *sync.WaitGroup) {
	// 'defer'is a magic word. It means: "Wait until this function is complete"
	// and then run this line right before ejecting
	// wg.Done() subtracts 1 from our WaitGroup counter
	defer wg.Done()

	fmt.Println("--> Sending ping to:", serverName)

	// We force the CPU to sleep for 2 seconds to simulate network lag.
	time.Sleep(2 * time.Second)

	fmt.Println("[OK] Response received from:", serverName)
}

func main() {
	// 2. THE INFRASTRUCTURE
	servers := []string{
		"Alpha-DB",
		"Beta-Cache",
		"Gamma-Web",
		"Delta-Auth",
	}

	// 3. THE CHECK COUNTER
	// We create a WaitGroup to keep track of how many background jobs are running.
	var wg sync.WaitGroup

	fmt.Println("--- INITIATING CONCURRENT SCAN ---")

	// 4. SPAWNING GOROUTINES
	for _, name := range servers {
		// We tell the WaitGroup: "Add 1 to the counter, a new job is starting"
		wg.Add(1)

		// The 'go' keyword fires this function off into the background.
		// It doesn't wait for the 2-second sleep. It instantly loops to the next server.
		go checkServerHealth(name, &wg)
	}

	// 5. THE BLOCKER
	// wg.Wait() pauses the main program right here.
	// It refuses to let the program move forward until the Waitgroup counter hits exactly 0.
	wg.Wait()

	fmt.Println("--- ALL SCANS COMPLETE ---")
}
