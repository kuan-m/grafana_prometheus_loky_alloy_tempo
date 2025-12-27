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
		var logLevel, logMessage string

		switch i % 3 {
		case 0:
			logLevel = "INFO"
			logMessage = "Application running normally"
		case 1:
			logLevel = "WARNING"
			logMessage = "Resource usage high"
		default:
			logLevel = "ERROR"
			logMessage = "Database connection lost"
		}

		component := components[rand.Intn(len(components))]
		timestamp := time.Now().UTC().Format(time.RFC3339)

		logEntry := fmt.Sprintf(
			"ts=%s level=%s app=myapp component=%s msg=%q\n",
			timestamp,
			logLevel,
			component,
			logMessage,
		)

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
