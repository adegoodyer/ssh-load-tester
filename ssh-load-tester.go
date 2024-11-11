package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	server := "your_server_ip_or_hostname"
	port := "22"
	username := "your_username"
	password := "wrong_password"

	attempts := 50

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

		// Delay between attempts to simulate real-world behavior
		time.Sleep(500 * time.Millisecond)
	}
}
