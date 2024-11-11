package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Configure server details
	server := "your_server_ip_or_hostname"
	port := "22"
	username := "your_username"
	password := "wrong_password" // Intentionally incorrect password

	// Set number of login attempts and rate per second
	attempts := 50
	ratePerSecond := 5 // Set rate of attempts per second

	// Calculate delay between attempts to match the rate per second
	delay := time.Second / time.Duration(ratePerSecond)

	// Run the login attempts
	for i := 0; i < attempts; i++ {
		// Start SSH connection
		clientConfig := &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

		// Attempt to connect
		conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", server, port), clientConfig)
		if err != nil {
			log.Printf("Attempt %d failed: %s\n", i+1, err)
		} else {
			conn.Close()
			log.Printf("Attempt %d succeeded\n", i+1)
		}

		// Delay between attempts to control the rate
		time.Sleep(delay)
	}
}
