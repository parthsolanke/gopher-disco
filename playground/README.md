## **Stage 1 – Warm-Up: Goroutines Basics**

### Tasks:

1. **Hello Goroutine**
   Write a program that:

  * Starts 5 goroutines, each printing `"Hello from Goroutine X"`.
  * The main function waits for them to finish without using `time.Sleep` (hint: `sync.WaitGroup`).

2. **Concurrent Countdown**

  * Have 3 goroutines each count down from 5 to 1, with a small random delay.
  * Observe the interleaving of prints.

3. **Parallel Square Calculator**

  * Given an array of numbers, spawn a goroutine for each number that computes its square.
  * Collect and print results.

---

## **Stage 2 – Channels Mastery**

Goal: Master unbuffered, buffered, and directional channels.

### Tasks:

4. **Ping-Pong** *(Unbuffered)*

  * Two goroutines play "ping-pong" by sending/receiving a string through a channel 10 times.

5. **Batch Logger** *(Buffered)*

  * Buffered channel stores log messages from multiple goroutines.
  * A single logger goroutine reads from it and prints messages.

6. **One-Way Streets** *(Directional)*

  * Implement a producer goroutine (send-only channel) that generates numbers 1–10.
  * Consumer goroutine (receive-only channel) sums them up.

---

## **Stage 3 – Select Statement Skills**

Goal: Handle multiple channels and timeouts.

### Tasks:

7. **Racing Goroutines**

  * Start two goroutines that each complete after a random time (1–3 sec).
  * Use `select` to print whichever finishes first.

8. **Timeout Fetcher**

  * Simulate fetching data from a service (sleep 1–5 sec).
  * If it takes more than 3 seconds, print `"Timeout"`.

9. **Multi-Source Data**

  * Simulate receiving weather updates from 3 sensors via 3 channels.
  * Use `select` in a loop to process whichever update arrives first.

---

## **Stage 4 – Coordination Patterns**

Goal: Learn common concurrency patterns in Go.

### Tasks:

10. **Worker Pool**

  * Implement a pool of N worker goroutines that process jobs from a channel.
  * When all jobs are done, stop the workers.

11. **Pipeline**

  * Stage 1: Generate numbers 1–20.
  * Stage 2: Filter even numbers.
  * Stage 3: Square them.
  * Each stage is its own goroutine, connected via channels.

12. **Fan-In**

  * Have two goroutines produce messages at different intervals.
  * Merge them into a single channel and print all messages.

13. **Fan-Out**

  * One producer sends jobs to multiple worker goroutines for parallel processing.

---

## **Stage 5 – Real-World Mini Projects**

Goal: Combine everything.

### Mini Projects:

14. **Concurrent Web Scraper**

  * Fetch 5 URLs concurrently.
  * Print the status code and time taken for each.

15. **Chat Room Simulation**

  * Multiple “users” (goroutines) send messages to a central chat room.
  * The chat room broadcasts messages to all connected users (fan-out).

16. **Stock Price Monitor**

  * Three simulated APIs send stock updates.
  * Use `select` to process updates in real-time and log them.

17. **Race Between Services** *(Like API fallbacks)*

  * Query two simulated services for the same data.
  * Use whichever responds first, cancel the slower one.

18. **Concurrent File Search**

  * Search for a keyword in multiple files concurrently.
  * Report matches as they are found.

---

## **Stage 6 – Extreme Challenges**

Goal: Push Go’s concurrency to the limit.

### Advanced:

19. **100,000 Goroutines Stress Test**

  * Start 100k goroutines doing small tasks.
  * Measure memory usage.

20. **Rate Limiter with Channels**

  * Limit function execution to 5 calls per second using a ticker channel.

21. **Deadlock Detector** *(Deliberate bug)*

  * Write a program that deadlocks.
  * Then fix it using correct channel/goroutine logic.

---

## Practice Approach

* **Start small** → don’t jump to projects until you master Stage 1–2.
* **No `time.Sleep` unless simulating work** → use `sync.WaitGroup` or channel signals for coordination.
* **Instrument your code** → print goroutine IDs or timestamps to see concurrency in action.
* **Refactor for clarity** → concurrency mistakes hide in messy code.
