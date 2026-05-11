package main

import "fmt"

// 1. THE BLUEPRINT
// 'type' tells Go we are inventing a new shape.
// 'struct' means it will be a collection of other variables.
type Server struct {
	Name      string
	IPAddress string
	Crashes   int
	IsOnline  bool
}

func main() {
	fmt.Println("--- PROVISIONING NEW INFRASTRUCTURE ---")

	// 2. BUILDING AN INSTANCE
	// We use our 'Server' blueprint to create a specific machine.
	// We use := to create and assign it to the variable 'frontendServer'.
	frontendServer := Server{
		Name:      "Web-Front-01",
		IPAddress: "192.168.1.50",
		Crashes:   0,
		IsOnline:  true, // true means it is running!
	}
	dataBaseServer := Server{
		Name:      "Customer-DB-01",
		IPAddress: "10.0.0.5",
		Crashes:   5,
		IsOnline:  false, // false means it is down!
	}

	// 3. ACCESSING DATA (Dot Notation)
	// We use a period '.' to look inside the struct and grab a specific field.
	fmt.Println("Booting:", frontendServer.Name)
	fmt.Println("IP:", frontendServer.IPAddress)

	fmt.Println("Booting:", dataBaseServer.Name)
	fmt.Println("IP:", dataBaseServer.IPAddress)

	// 4. THE SRE SCENARIO
	// A power blip hits the data center. The server goes down.
	fmt.Println("... POWER ANOMALY DETECTED ...")

	// We use dot notation to modify the internal state of our struct.
	frontendServer.IsOnline = false
	frontendServer.Crashes = frontendServer.Crashes + 1

	dataBaseServer.IsOnline = false
	dataBaseServer.Crashes = dataBaseServer.Crashes + 1

	// 5. OUTPUT NEW STATE
	fmt.Println("Server Online:", frontendServer.IsOnline)
	fmt.Println("Total Crashes:", frontendServer.Crashes)
	fmt.Println("Server Online:", dataBaseServer.IsOnline)
	fmt.Println("Total Crashes:", dataBaseServer.Crashes)
	fmt.Printf("ALERT: %s at %s is currently offline!\n", frontendServer.Name, frontendServer.IPAddress)
	fmt.Printf("ALERT: %s at %s is currently offline!\n", dataBaseServer.Name, dataBaseServer.IPAddress)

}
