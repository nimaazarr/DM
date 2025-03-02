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

	
	fmt.Println("🚀 Starting concurrent downloads...")

	// Initialize the worker pool
	workerPool := download.NewWorkerPool(conf.MaxConcurrentDownloads)
	workerPool.Start()

	// Test URLs (you can modify this list)
	testURLs := []string{
		utils.TestURLs["10MB_OVH"],
		utils.TestURLs["100MB_OVH"],
		// utils.TestURLs["10MB_OVH"],
		// utils.TestURLs["1MB"],
	}

	// Queue downloads
	for _, url := range testURLs {
		workerPool.AddJob(url, conf.DownloadDirectory)
	}

	// Wait for all downloads to complete
	workerPool.Wait()

	fmt.Println("🎉 All downloads completed!")
}