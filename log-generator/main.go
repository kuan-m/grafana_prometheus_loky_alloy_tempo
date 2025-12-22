package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	logFilePath = "/var/log/loki_udemy.log"
)

var (
	components = []string{"database", "backend"}
	logLevels  = []string{"INFO", "WARNING", "ERROR"}
)

func main() {
	// Open log file for appending
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Starting log generation...")

	for i := 0; i < 10; i++ {
		// Determine log level based on iteration
		var logLevel string
		var logMessage string

		switch i % 3 {
		case 0:
			logLevel = "INFO"
			logMessage = "Information: Application running normally"
		case 1:
			logLevel = "WARNING"
			logMessage = "Warning: Resource usage high"
		default:
			logLevel = "ERROR"
			logMessage = "Critical error: Database connection lost"
		}

		// Randomly select component
		component := components[rand.Intn(len(components))]

		// Format log entry
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logEntry := fmt.Sprintf("%s level=%s app=myapp component=%s %s\n",
			timestamp, logLevel, component, logMessage)

		// Write to file
		if _, err := file.WriteString(logEntry); err != nil {
			log.Printf("Failed to write log entry: %v", err)
			continue
		}

		fmt.Printf("Generated log: level=%s component=%s\n", logLevel, component)

		// Sleep for 1 second between entries
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Log generation completed!")
}

