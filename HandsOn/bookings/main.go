package main

import (
	"fmt"
	"sync"
)

type Ticket struct {
	isSold bool
	id     int
}

type Concert struct {
	tickets      []Ticket
	totalTickets int
	mu           sync.RWMutex
}

func (c *Concert) book(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if id < 0 || id >= len(c.tickets) {
		fmt.Printf("nnvalid ticket ID %d\n", id)
		return
	}
	if c.tickets[id].isSold {
		fmt.Printf("ticket %d already sold\n", id)
		return
	}

	c.tickets[id].isSold = true
	c.totalTickets--
	fmt.Printf("booked ticket %d successfully, remaining tickets: %d\n", id, c.totalTickets)
}

func worker(id int, jobs <-chan int, c *Concert, wg *sync.WaitGroup) {
	defer wg.Done()
	for ticketID := range jobs {
		fmt.Printf("worker %d processing booking for ticket %d\n", id, ticketID)
		c.book(ticketID)
	}
}

func main() {
	numTickets := 1000000
	myTickets := make([]Ticket, numTickets)
	for i := range numTickets {
		myTickets[i] = Ticket{id: i}
	}

	c := Concert{
		tickets:      myTickets,
		totalTickets: numTickets,
	}

	const numWorkers = 1000
	jobs := make(chan int, 20)
	var wg sync.WaitGroup

	// init workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &c, &wg)
	}

	// create jobs
	for i := range numTickets {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	fmt.Println("booking requests processed.")

}
