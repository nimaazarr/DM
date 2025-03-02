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

	
	fmt.Println("🚀 Starting Download Manager with Scheduling...")

	// Initialize worker pool, queue, and scheduler
	workerPool := download.NewWorkerPool(conf.MaxConcurrentDownloads)
	workerPool.Start()
	queue := download.NewDownloadQueue(workerPool)
	scheduler := download.NewScheduler(queue)

	// // Test URLs (you can modify this list)
	// testURLs := []string{
	// 	utils.TestURLs["10MB_OVH"],
	// 	utils.TestURLs["10MB_OVH"],
	// 	// utils.TestURLs["10MB_OVH"],
	// 	// utils.TestURLs["1MB"],
	// }

	// for _, url := range testURLs {
	// 	queue.AddDownload(url, conf.DownloadDirectory)
	// }

	// // Start downloads
	// for range testURLs {
	// 	queue.StartNextDownload()
	// }




	// Add immediate downloads
	queue.AddDownload(utils.TestURLs["10MB_OVH"], conf.DownloadDirectory)
	queue.StartNextDownload()

	// Schedule a download 10 seconds in the future
	startTime := time.Now().Add(10 * time.Second)
	scheduler.ScheduleDownload(utils.TestURLs["10MB_OVH"], conf.DownloadDirectory, startTime)

	// List scheduled downloads
	scheduler.ListScheduledDownloads()

	// Wait for all downloads to complete
	workerPool.Wait()

	fmt.Println("🎉 All downloads (including scheduled ones) completed!")

	// Wait for all downloads to complete
	workerPool.Wait()

	fmt.Println("🎉 All downloads completed with speed control!")
}