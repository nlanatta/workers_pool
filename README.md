# Example of main
```go
package main

import "time"

type CustomJob struct {
	Id     int
	Result chan int
}

func (j *CustomJob) Process() error {
	time.Sleep(1 * time.Second)
	return nil
}

func (j *CustomJob) ID() int {
	return j.Id
}

func main() {
	wp := NewWorkerPool(3, 5)

	wp.Start()

	for i := 1; i <= 50; i++ {
		task := &CustomJob{Id: i}
		wp.AddJob(task)
	}

	time.Sleep(10 * time.Second)
	wp.Stop()
}
```

Run tests
```
go test -v -race ./...
```