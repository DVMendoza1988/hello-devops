package main

import "fmt"

func main() {
	fmt.Println("--- FETCHING INCIDENT LOG ---")

	// 1. THE SLICE
	// '[]string' strictly means "This is an ordered list containing ONLY text."
	// We populate it with three initial events.
	var incidentLog []string = []string{
		"08:00 AM - Server booted normally",
		"09:15 AM - High memory usage warning",
		"09:20 AM - Auto-scaling initiated",
	}

	// 2. APPENDING
	// An SRE's job is constantly updating state.
	// 'append' takes our existing list, adds a new item to the end,
	// and saves the new, longer list back into the incidentLog variable.
	incidentLog = append(incidentLog, "10:05 AM - Database connection lost")

	// 3. LOOPING A SLICE
	// When we range over a Slice, it hands us two things:
	// The 'index' (its number in line) and the 'event' (the text data).
	for index, event := range incidentLog {
		// Notice how we combine text, an integer, and text again.
		fmt.Println("Event ID", index, "->", event)
	}
}
