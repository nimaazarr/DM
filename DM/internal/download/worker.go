package download

import (
	"fmt"
	"log"
	"sync"
	// "time"

	// "DM/DM/internal/config"
)

// DownloadJob represents a download task.
type DownloadJob struct {
	URL      string
	DestPath string
}

// WorkerPool manages concurrent download workers.
type WorkerPool struct {
	MaxWorkers int
	JobQueue   chan DownloadJob
	WG         sync.WaitGroup
}

// NewWorkerPool initializes a new worker pool.
func NewWorkerPool(maxWorkers int) *WorkerPool {
	return &WorkerPool{
		MaxWorkers: maxWorkers,
		JobQueue:   make(chan DownloadJob, maxWorkers),
	}
}

// Start initializes worker goroutines.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.MaxWorkers; i++ {
		go wp.worker(i)
	}
}

// worker executes download jobs.
func (wp *WorkerPool) worker(workerID int) {
	for job := range wp.JobQueue {
		fmt.Printf("🔄 Worker %d: Downloading %s\n", workerID, job.URL)

		// Perform file download
		_, err := DownloadFile(job.URL, job.DestPath)
		if err != nil {
			log.Printf("❌ Worker %d: Failed to download %s: %v\n", workerID, job.URL, err)
		} else {
			fmt.Printf("✅ Worker %d: Finished %s\n", workerID, job.URL)
		}
		wp.WG.Done()
	}
}

// AddJob queues a new download job.
func (wp *WorkerPool) AddJob(url string, destPath string) {
	wp.WG.Add(1)
	wp.JobQueue <- DownloadJob{URL: url, DestPath: destPath}
}

// Wait waits for all jobs to complete.
func (wp *WorkerPool) Wait() {
	wp.WG.Wait()
	close(wp.JobQueue) // Close the channel when done
}