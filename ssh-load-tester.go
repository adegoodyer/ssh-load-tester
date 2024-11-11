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

	// Set number of login attempts and rate per second
	attempts := 50
	ratePerSecond := 5 // Set rate of attempts per second

	// Calculate the interval for the specified rate per second
	interval := time.Second / time.Duration(ratePerSecond)

	// Create WaitGroup for managing parallel attempts
	var wg sync.WaitGroup

	// Run the login attempts in parallel
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go attemptLogin(server, port, username, password, &wg)

		// Wait for the next attempt
		time.Sleep(interval)
	}

	// Wait for all goroutines to complete
	wg.Wait()
}
