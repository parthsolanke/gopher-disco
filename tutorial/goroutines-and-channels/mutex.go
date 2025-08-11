package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string, n int, wg *sync.WaitGroup) {
	for range n {
		c.mu.Lock()
		c.counters[name]++
		c.mu.Unlock()
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := Container{counters: map[string]int{"a": 0, "b": 0}}

	wg.Add(3)
	c.inc("a", 200, &wg)
	c.inc("a", 200, &wg)
	c.inc("b", 200, &wg)
	wg.Wait()

	fmt.Println(c.counters)
}
