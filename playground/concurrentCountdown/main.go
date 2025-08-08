package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func countdown(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 3; i > 0; i-- {
		fmt.Printf("id: %d - count: %d\n", id, i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	n := 3

	wg.Add(n)
	for i := range 3 {
		go countdown(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("Finish countdown")
}
