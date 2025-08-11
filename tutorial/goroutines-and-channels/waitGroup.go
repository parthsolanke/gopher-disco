package main

import (
	"fmt"
	"sync"
	"time"
)

func goodWorker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("myWorker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("myWorker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	numJobs := 5
	numWorkers := 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go goodWorker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 0; j < numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish, then close results
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for r := range results {
		fmt.Println("Result:", r)
	}
}
