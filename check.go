package main

import (
	"fmt"
	"net"
	"time"
)

// Check function attempts to establish a TCP connection to a destination and port,
// and returns a string indicating whether the destination is reachable or not.
func Check(destination string, port string) string {
	// Construct the full address by combining destination and port
	address := destination + ":" + port

	// Set timeout duration for the connection attempt
	timeout := time.Duration(5 * time.Second)

	// Attempt to establish a TCP connection with timeout
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string // Initialize variable to hold status message

	// Check if there was an error during connection attempt
	if err != nil {
		// Format status message indicating destination is unreachable
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		// Format status message indicating destination is reachable
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination, conn.LocalAddr(), conn.RemoteAddr())
	}

	return status // Return the status message indicating UP or DOWN
}
