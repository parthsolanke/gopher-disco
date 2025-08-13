## **Go Concurrency Practice Problems**

### **1. Coffee Shop Order System** *(Goroutines + Unbuffered Channels)*

**Scenario:**
You’re building a coffee shop simulator. Each customer places an order, and the barista prepares it.
**Requirements:**

* Each order is handled in its own goroutine.
* Orders are sent through an **unbuffered channel** from customers to barista.
* The barista confirms each completed drink back via another channel.
* Stop the simulation after serving 5 customers.

---

### **2. Log Processing System** *(Buffered Channels + Goroutines)*

**Scenario:**
A web server produces log messages faster than they can be processed.
**Requirements:**

* Use a **buffered channel** to store logs temporarily.
* A single processor goroutine consumes logs from the channel and prints them.
* Simulate producers generating logs at random intervals, and consumer processing slower.
* Stop after 20 logs.

---

### **3. Stock Price Aggregator** *(Select + Timeouts)*

**Scenario:**
You need to get a stock’s price from **3 different APIs** (simulated by goroutines with random delays).
**Requirements:**

* Use **`select`** to listen for the first API to respond.
* If no API responds within 2 seconds, print “Timeout”.
* Cancel the slower API calls once you have the result.

---

### **4. Image Downloader** *(WaitGroup)*

**Scenario:**
You have a list of image URLs to download.
**Requirements:**

* Launch a goroutine per download.
* Use **`sync.WaitGroup`** to wait for all downloads to finish.
* Print “All downloads completed” at the end.
* Simulate downloads with `time.Sleep` for different durations.

---

### **5. Bank Account System** *(Mutex for Data Safety)*

**Scenario:**
A bank account is shared by multiple users depositing and withdrawing money.
**Requirements:**

* Use **`sync.Mutex`** to ensure only one goroutine updates the balance at a time.
* Simulate multiple deposits and withdrawals running concurrently.
* Print the final balance after all operations complete.

---

### **6. Read-Mostly Cache** *(RWMutex)*

**Scenario:**
You have a shared product catalog that is **read** frequently but **updated** rarely.
**Requirements:**

* Use **`sync.RWMutex`** so multiple goroutines can read simultaneously, but writes block reads and other writes.
* Simulate:

   * Many readers fetching product prices.
   * An occasional writer updating a price.
* Ensure data remains consistent.

---

### **7. Video Encoding Farm** *(Worker Pool + WaitGroup)*

**Scenario:**
You run a video encoding service with many jobs but want to limit concurrent encodings to 3 at a time.
**Requirements:**

* Implement a **worker pool** with a fixed number of workers (goroutines).
* Use a **job channel** to queue videos.
* Each worker processes jobs and sends results to a results channel.
* Use **WaitGroup** to wait until all videos are processed.

---

### **8. Real-Time Chat Server** *(Context + Multiple Concepts Combined)*

**Scenario:**
Build a mini chat server where multiple clients can send messages until a **context timeout** stops the server.
**Requirements:**

* Use **goroutines** for each client connection.
* Use a **channel** for broadcasting messages.
* Use **select** to handle multiple message channels.
* Use **`context.WithTimeout`** to automatically shut down after 10 seconds.
* Gracefully stop all goroutines when context is done.
