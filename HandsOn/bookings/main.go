package main

import (
	"fmt"
	"math/rand"
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
	fmt.Printf("booked ticket %d successfully.\n", id)
}

func worker(id int, jobs <-chan int, c *Concert, wg *sync.WaitGroup) {
	defer wg.Done()
	for ticketID := range jobs {
		fmt.Printf("worker %d processing booking for ticket %d\n", id, ticketID)
		c.book(ticketID)
	}
}

func main() {
	numTickets := 100
	myTickets := make([]Ticket, numTickets)
	for i := range numTickets {
		myTickets[i] = Ticket{id: i}
	}

	c := Concert{
		tickets:      myTickets,
		totalTickets: numTickets,
	}

	const numWorkers = 5
	jobs := make(chan int, 20)
	var wg sync.WaitGroup

	// init workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &c, &wg)
	}

	// create jobs
	for i := 0; i < 80; i++ {
		jobs <- rand.Intn(numTickets)
	}
	close(jobs)

	wg.Wait()
	fmt.Println("booking requests processed.")

}
