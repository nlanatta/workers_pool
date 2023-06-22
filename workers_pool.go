package main

import (
	"fmt"
	"sync"
)

type WorkersPool struct {
	numWorkers int
	jobQueue   chan Job
	wg         *sync.WaitGroup
	quitCh     chan struct{}
}

func NewWorkerPool(numWorkers, queueSize int) *WorkersPool {
	return &WorkersPool{
		numWorkers: numWorkers,
		jobQueue:   make(chan Job, queueSize),
		quitCh:     make(chan struct{}),
		wg:         &sync.WaitGroup{},
	}
}

func (wp *WorkersPool) AddJob(task Job) {
	wp.wg.Add(1)
	wp.jobQueue <- task
}

func (wp *WorkersPool) Start() {
	for i := 1; i <= wp.numWorkers; i++ {
		go wp.worker(i)
	}

	fmt.Println("Worker pool started")
}

func (wp *WorkersPool) Wait() {
	wp.wg.Wait()
}

func (wp *WorkersPool) Stop() {
	fmt.Println("Worker pool stopped")
	close(wp.jobQueue)
	close(wp.quitCh)
}

func (wp *WorkersPool) worker(id int) {
	for {
		select {
		case task, ok := <-wp.jobQueue:
			if !ok {
				fmt.Printf("Worker %d shutting down\n", id)
				return
			}
			err := task.Process()
			if err != nil {
				fmt.Printf("Worker %d, ERROR processing task: %d\n", id, task.ID())
			}
			fmt.Printf("Worker %d processing task: %d\n", id, task.ID())
			wp.wg.Done()
		case <-wp.quitCh:
			fmt.Printf("Worker %d stopped\n", id)
			wp.Stop()
			return
		}
	}
}
