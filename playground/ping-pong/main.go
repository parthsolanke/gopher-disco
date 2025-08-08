package main

import (
	"fmt"
	"sync"
)

func ping(pingCh chan<- string, pongCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		pingCh <- fmt.Sprintf("ping %d", i)
		fmt.Println(<-pongCh)
	}
}

func pong(pingCh <-chan string, pongCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		msg := <-pingCh
		fmt.Println(msg)
		pongCh <- fmt.Sprintf("pong %d", i) // send pong
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	pingCh := make(chan string)
	pongCh := make(chan string)

	go ping(pingCh, pongCh, &wg)
	go pong(pingCh, pongCh, &wg)

	wg.Wait()
}
