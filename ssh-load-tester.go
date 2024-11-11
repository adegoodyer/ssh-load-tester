package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

func attemptLogin(server, port, username, password string, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", server, port), clientConfig)
	if err != nil {
		log.Println("Login failed:", err)
		return
	}
	conn.Close()
	log.Println("Login succeeded")
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

	// Create WaitGroup for managing parallel attempts
	var wg sync.WaitGroup

	// Run the login attempts for the specified duration
	for time.Now().Before(endTime) {
		wg.Add(1)
		go attemptLogin(server, port, username, password, &wg)

		// Wait for the next attempt to control the rate
		time.Sleep(interval)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
