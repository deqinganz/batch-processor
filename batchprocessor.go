package batchprocessor

import (
	"fmt"
	. "github.com/deqinganz/batching-api/api"
	"log"
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
