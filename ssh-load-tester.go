package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func attemptLogin(server, port, username, password string, wg *sync.WaitGroup, attemptNumber int) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Log the attempt number
	log.Printf("Starting attempt %d...\n", attemptNumber)

	// Attempt to connect
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", server, port), clientConfig)
	if err != nil {
		log.Printf("Attempt %d failed: %s\n", attemptNumber, err)
		return
	}
	conn.Close()
	log.Printf("Attempt %d succeeded\n", attemptNumber)
}

func main() {
	// Configure server details
	server := "your_server_ip_or_hostname"
	port := "22"
	username := "your_username"
	password := "wrong_password" // Intentionally incorrect password

	// Set duration for the test and rate per second
	duration := 10 * time.Second // Run for 10 seconds
	ratePerSecond := 5           // Set rate of attempts per second

	// Calculate interval between attempts to match the rate per second
	interval := time.Second / time.Duration(ratePerSecond)
	endTime := time.Now().Add(duration)

	// Display test summary
	fmt.Println("----- Test Summary -----")
	fmt.Printf("Server: %s\n", server)
	fmt.Printf("Port: %s\n", port)
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Test Duration: %s\n", duration)
	fmt.Printf("Rate of Attempts: %d per second\n", ratePerSecond)
	fmt.Println("------------------------")

	// Create WaitGroup for managing parallel attempts
	var wg sync.WaitGroup

	// Run the login attempts for the specified duration
	attemptNumber := 1
	for time.Now().Before(endTime) {
		wg.Add(1)
		go attemptLogin(server, port, username, password, &wg, attemptNumber)

		// Increment the attempt number and wait for the next attempt to control the rate
		attemptNumber++
		time.Sleep(interval)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Display final message after all attempts are complete
	fmt.Println("----- Test Completed -----")
	fmt.Printf("Total Attempts Made: %d\n", attemptNumber-1)
	fmt.Println("All login attempts finished.")
}
