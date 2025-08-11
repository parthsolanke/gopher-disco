package main

import (
	"fmt"
	"sync"
)

// RLock() — Acquire a read lock.
//Many goroutines can hold the read lock at the same time.
//But if anyone has a read lock, no one can acquire the write lock until all readers are done.

//Lock() — Acquire a write lock.
//Only one goroutine can hold it.
//While a write lock is held, no one can acquire a read lock.

// gist: many readers, one writer, never both at the same time.

type SafeCounter struct {
	mu    sync.RWMutex
	count int
}

func (s *SafeCounter) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	s.mu.Lock() // full lock while writing
	defer s.mu.Unlock()
	s.count++
}

func (s *SafeCounter) value(wg *sync.WaitGroup) int {
	defer wg.Done()

	s.mu.RLock() // only reader lock
	defer s.mu.RUnlock()
	return s.count
}

func main() {
	var wg sync.WaitGroup
	sc := SafeCounter{}

	// writers
	for range 5 {
		wg.Add(1)
		go sc.inc(&wg)
	}

	// readers
	for i := range 5 {
		wg.Add(1)
		go func(id int) {
			val := sc.value(&wg)
			fmt.Println("Reader", id, "saw value:", val)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final count: ", sc.count)
}
