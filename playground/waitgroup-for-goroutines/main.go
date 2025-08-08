package main

import (
	"fmt"
	"sync"
)

func hello(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("hello from goroutine: %d\n", id)
}

func main() {
	var wg sync.WaitGroup
	n := 5
	wg.Add(n)

	for i := range n {
		go hello(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("\nFinished all goroutines")
}
