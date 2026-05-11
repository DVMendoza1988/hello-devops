package main

import "fmt"

type Server struct {
	Name      string
	IPAddress string
	IsOnline  bool
	Crashes   int
}

// 1. THE METHOD
// Notice the special receiver block: (s *Server) BEFORE the function name.
// This physicall attaches the 'Reboot' function to the Server blueprint.
// * tells Go: "Do not make a copy. Use the Pointer to the actual memory."

func (s *Server) Reboot() {
	fmt.Printf("\n--- INITIATING REBOOT SEQUENCE FOR %s ---\n", s.Name)

	// Because we used a pointer (*), this changes the REAL server.
	s.IsOnline = true
	s.Crashes = 0

	fmt.Printf("SYSTEM MESSAGE: Boot successful.")
}

func main() {
	// 2. PROVISION A DEAD SERVER
	targetServer := Server{
		Name:      "Payment-Gateway-01",
		IPAddress: "172.16.254.1",
		IsOnline:  false,
		Crashes:   14,
	}
	// 3. CHECK INITIAL STATE
	fmt.Println("Initial State ->", targetServer.Name, "Online:", targetServer.IsOnline)

	// 4. PRESS THE REBOOT BUTTON
	// We use dot notation to trigger the method attached to this specific struct.
	targetServer.Reboot()

	// 5. VERIFY RECOVERY
	// If we didn't use the '*' pointer in step 1, this would still print 'false'!
	fmt.Printf("\nFinal State -> %s Online: %t\n Crashes: %d\n", targetServer.Name, targetServer.IsOnline, targetServer.Crashes)
}
