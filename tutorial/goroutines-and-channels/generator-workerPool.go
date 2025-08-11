package main

import (
	"fmt"
	"sync"
	"time"
)

// jobGenerator produces jobs and closes the channel when done
func jobGenerator(numJobs int) <-chan int {
	out := make(chan int)
	go func() {
		for j := 1; j <= numJobs; j++ {
			out <- j
		}
		close(out)
	}()
	return out
}

func workerGenerator(id int, jobs <-chan int, wg *sync.WaitGroup) <-chan int {
	results := make(chan int)
	wg.Add(1) // we have one more worker goroutine
	go func() {
		defer wg.Done() // mark this worker as done when it exits
		for j := range jobs {
			fmt.Println("workerGenerator", id, "started  job", j)
			time.Sleep(time.Second)
			fmt.Println("workerGenerator", id, "finished job", j)
			results <- j * 2
		}
		close(results)
	}()
	return results
}

func main() {
	numJobs := 5
	numWorkers := 3

	var wg sync.WaitGroup

	jobs := jobGenerator(numJobs)

	// Create workers
	var workerResults []<-chan int
	for w := 1; w <= numWorkers; w++ {
		workerResults = append(workerResults, workerGenerator(w, jobs, &wg))
	}

	// Wait for all workers in a separate goroutine, so we can close the merged results channel later
	go func() {
		wg.Wait()
		fmt.Println("All workers finished")
	}()

	// Collect results
	for _, results := range workerResults {
		for r := range results {
			_ = r // process r here
		}
	}
}
