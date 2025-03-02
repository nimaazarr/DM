package main

import (
	"fmt"
	"log"

	"DM/internal/config"
)

func main() {
	// Load config from file or environment variables
	conf, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Print loaded configuration
	config.PrintConfig(conf)

	fmt.Println("🚀 Download Manager Starting...")
}