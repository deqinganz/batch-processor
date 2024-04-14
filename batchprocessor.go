package batchprocessor

import (
	"fmt"
	. "github.com/deqinganz/batching-api/api"
	"log"
	"time"
)

type BatchProcessor struct{}

// Process simply changes all job status to SUBMITTED and logs the jobs that are processed
func (b *BatchProcessor) Process(jobs []Job) {
	for i := range jobs {
		jobs[i].Status = SUBMITTED
	}

	var output string
	for _, job := range jobs {
		output = output + fmt.Sprintf("[%s %s \"%s\"] ", job.Id.String()[:8], job.Type, job.Name)
	}
	log.Printf("Processed jobs: %s", output)
}

// ProcessAndSleep is a wrapper function that calls Process and then sleeps for a given duration
func (b *BatchProcessor) ProcessAndSleep(jobs []Job, duration int) {
	b.Process(jobs)
	log.Printf("Sleeping for %d seconds", duration)
	time.Sleep(time.Duration(duration) * time.Second)
}
