package main

import (
	"fmt"
	"log"

	"DM/DM/internal/config"
	"DM/DM/internal/download"
	"DM/DM/internal/utils" // Import the test URLs
)

func main() {
	// Load configuration
	conf, err := config.LoadConfig("internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Choose a test URL
	testURL := utils.TestURLs["10MB_OVH"] // Select desired file size

	fmt.Println("Starting download:", testURL)

	filePath, err := download.DownloadFile(testURL, conf.DownloadDirectory)
	if err != nil {
		log.Fatalf("Download failed: %v", err)
	}

	fmt.Println("✅ Download completed:", filePath)
}