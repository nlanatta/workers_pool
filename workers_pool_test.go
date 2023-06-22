package main

import (
	"sync"
	"testing"
)

var counter = 0

type TestJob struct {
	Id  int
	mut sync.Mutex
}

func (j *TestJob) Process() error {
	j.mut.Lock()
	defer j.mut.Unlock()
	counter++
	return nil
}

func (j *TestJob) ID() int {
	return j.Id
}

func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(3, 1)
	expected := 10

	wp.Start()

	for i := 1; i <= expected; i++ {
		task := &TestJob{Id: i}
		wp.AddJob(task)
	}

	wp.Wait()
	result := counter

	if result != expected {
		t.Fatalf("expected 10, got %d", result)
	}
}
