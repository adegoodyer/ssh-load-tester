package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

// Custom log format to show time with milliseconds at the start of each log entry
func logWithTimestamp(msg string, args ...interface{}) {
	// Format the current time with milliseconds
	currentTime := time.Now().Format("15:04:05.000")
	fmt.Printf("%s %s\n", currentTime, fmt.Sprintf(msg, args...))
}

func attemptLogin(host, port, username, password string, wg *sync.WaitGroup, attemptNumber int, interval time.Duration) {
	defer wg.Done() // Decrement the counter when the goroutine completes

	clientConfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Log the attempt number with milliseconds using the custom log function
	logWithTimestamp("Starting attempt %d with interval %v...", attemptNumber, interval)

	// Attempt to connect
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", host, port), clientConfig)
	if err != nil {
		logWithTimestamp("Attempt %d failed: %s", attemptNumber, err)
		return
	}
	conn.Close()
	logWithTimestamp("Attempt %d succeeded", attemptNumber)
}

func main() {
	// Define short flags for host, port, username, password, duration, and rate per second
	host := flag.String("h", "", "host IP or hostname (required)")
	port := flag.String("p", "22", "SSH port")
	username := flag.String("U", "ssh-load-tester", "username for SSH login")
	password := flag.String("P", "1nV@l!dP@ss", "password for SSH login")
	duration := flag.Duration("d", 10*time.Second, "test duration")
	ratePerSecond := flag.Int("r", 5, "rate of attempts per second")

	// Parse command-line flags
	flag.Parse()

	// Check if -h (host) is provided; if not, show usage and exit
	if *host == "" {
		fmt.Println("Error: -h flag (host) is required.")
		flag.Usage()
		os.Exit(1)
	}

	// Calculate interval between attempts to match the rate per second
	interval := time.Second / time.Duration(*ratePerSecond)
	endTime := time.Now().Add(*duration)

	// Display test summary including interval
	fmt.Println("----- Test Summary -----")
	fmt.Printf("Host: %s\n", *host)
	fmt.Printf("Port: %s\n", *port)
	fmt.Printf("Username: %s\n", *username)
	fmt.Printf("Test Duration: %s\n", *duration)
	fmt.Printf("Rate of Attempts: %d per second\n", *ratePerSecond)
	fmt.Printf("Interval between Attempts: %v\n", interval)
	fmt.Println("------------------------")

	// Create WaitGroup for managing parallel attempts
	var wg sync.WaitGroup

	// Run the login attempts for the specified duration
	attemptNumber := 1
	for time.Now().Before(endTime) {
		wg.Add(1)
		go attemptLogin(*host, *port, *username, *password, &wg, attemptNumber, interval)

		// Increment the attempt number and wait for the next attempt to control the rate
		attemptNumber++
		time.Sleep(interval)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Display final message after all attempts are complete
	logWithTimestamp("----- Test Completed -----")
	fmt.Printf("Total Attempts Made: %d\n", attemptNumber-1)
	fmt.Println("All login attempts finished.")
}
