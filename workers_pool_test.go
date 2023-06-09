package main

import (
	"testing"
)

type TestJob struct {
	Id     int
	Result chan int
}

func (j *TestJob) Process() error {
	j.Result <- 1
	return nil
}

func (j *TestJob) ID() int {
	return j.Id
}

func TestNewWorkerPool(t *testing.T) {
	wp := NewWorkerPool(3, 1)
	expected := 10

	wp.Start()

	c := make(chan int, expected)

	for i := 1; i <= expected; i++ {
		task := &TestJob{Id: i, Result: c}
		wp.AddJob(task)
	}

	result := 0
	for v := range c {
		result += v
		if result == expected {
			close(c)
		}
	}

	if result != expected {
		t.Fatalf("expected 10, got %d", result)
	}
}
