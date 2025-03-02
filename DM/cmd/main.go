package main

import (
	"fmt"
	"log"
	"time"

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

	// Initialize worker pool and download queue
	workerPool := download.NewWorkerPool(conf.MaxConcurrentDownloads)
	workerPool.Start()
	queue := download.NewDownloadQueue(workerPool)

	// Test URLs (you can modify this list)
	testURLs := []string{
		utils.TestURLs["10MB_OVH"],
		utils.TestURLs["100MB_OVH"],
		// utils.TestURLs["10MB_OVH"],
		// utils.TestURLs["1MB"],
	}

	for _, url := range testURLs {
		queue.AddDownload(url, conf.DownloadDirectory)
	}

	// Start downloads
	for range testURLs {
		queue.StartNextDownload()
	}

	// List active downloads
	queue.ListDownloads()

	// Simulate a pause/resume cycle
	time.Sleep(2 * time.Second)
	queue.PauseDownload(testURLs[1])

	time.Sleep(2 * time.Second)
	queue.ResumeDownload(testURLs[1])

	// Wait for all downloads to complete
	workerPool.Wait()

	fmt.Println("🎉 All downloads completed!")
	queue.ListDownloads()
}