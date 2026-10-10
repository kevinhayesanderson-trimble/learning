package jobprocessor

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type JobStatus int

const (
	Submitted JobStatus = iota
	Created
	Finished 	
)

type Job struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
}

type JobResponse struct{
	ID        string    `json:"id"`
	Payload   string    `json:"payload"`
	CreatedAt time.Time `json:"created_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	JobStatus 
}

type JobProcessor struct {
	jobQueue   chan Job
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	mu         sync.RWMutex
	isStopping bool
}

func NewJobProcessor(queueSize int) *JobProcessor {
	ctx, cancel := context.WithCancel(context.Background())
	jobQueue := make(chan Job, queueSize)
	processor := &JobProcessor{
		jobQueue:   jobQueue,
		ctx:        ctx,
		cancel:     cancel,
		isStopping: false,
	}
	return processor
}

func (p *JobProcessor) Submit(job Job) error {
	p.mu.RLock()
	defer p.mu.RUnlock()
	if p.isStopping {
		return errors.New("process is stopping, unable to submit this job")
	}
	select {
	case p.jobQueue <- job:
		return nil
	default:
		return fmt.Errorf("Job queue full")
	}
}

func (p *JobProcessor) worker() {
	defer p.wg.Done()
	for {
		select {
		case job, ok := <-p.jobQueue:
			if !ok {
				return
			}
			fmt.Println(job.ID, job.CreatedAt, job.Payload)	
		case <-p.ctx.Done():
			return
		}
	}
}

func (p *JobProcessor) StartWorkers(count int) {
	for range count {
		p.wg.Add(1)
		go p.worker()
	}
}

func (p *JobProcessor) Stop(){
	p.mu.Lock()
	if p.isStopping{
		p.mu.Unlock()
		return
	}
	p.isStopping = true
	p.cancel()
	close(p.jobQueue)
	p.mu.Unlock()
	p.wg.Wait()
}
