## 🧠 MASTERING GO FUNDAMENTALS — PRACTICE SET

### 🟢 LEVEL 1: Foundations & Flow

---

### **1. Student Grader**

🧩 **Concepts**: Input/output, variables, conditionals, switch, formatting
📌 **Task**:

* Input a student's name and marks for 3 subjects.
* Compute average and assign grades using `switch`.
* Print a summary like:

  ```
  Name: Raj | Avg: 72.33 | Grade: B
  ```

🔁 Bonus: Accept input in a loop till user types `exit`.

---

### **2. Prime Number Checker**

🧩 **Concepts**: `for` loops, conditionals, return, function design
📌 **Task**:

* Write a function `isPrime(n int) bool`
* Take user input for a number, print whether it's prime.
* Use it inside a loop that checks primes from 1 to 100.

⚙️ Optional: Add a recursive variant.

---

### **3. Shopping Cart Calculator**

🧩 **Concepts**: variables, loops, `range`, function with multiple return values
📌 **Task**:

* Given a slice of prices, write a function:

  ```go
  func totalCost(prices []float64) (float64, int)
  ```
* It should return total cost and number of items.
* Use `range` to loop over the slice.

---

## 🟡 LEVEL 2: Functions & Flow Composition

---

### **4. Password Strength Evaluator**

🧩 **Concepts**: string operators, if-else, functions, multiple return values
📌 **Task**:

* Write a function:

  ```go
  func checkStrength(pwd string) (int, string)
  ```
* Score based on:

    * Length
    * Presence of digits/symbols/uppercase
* Return score and comment ("Weak", "Strong", etc.)

---

### **5. Number Series Generator**

🧩 **Concepts**: variadic function, recursion, slice ops
📌 **Task**:

* Write `generateSeries(start int, steps ...int) []int`
* It should return a series where each number is:

  ```go
  next = previous + step[i % len(steps)]
  ```
* Example:

  ```go
  generateSeries(1, 2, 3) → [1, 3, 6, 8, 11, ...]
  ```

---

### **6. Simple CSV Parser**

🧩 **Concepts**: string splitting, `range`, `defer`, file I/O (optional)
📌 **Task**:

* Parse a simulated CSV like:

  ```go
  "Alice,89\nBob,76\nCharlie,92"
  ```
* Output:

  ```go
  Alice got 89
  ```
* Defer closing file (optional: read from file using `os.Open`)

---

## 🔴 LEVEL 3: Combined Concepts & Design Thinking

---

### **7. Bank Account System (In-Memory)**

🧩 **Concepts**: struct (when you're ready), switch, functions, defer
📌 **Task**:

* Create a CLI-like system:

    * `deposit(amount)`
    * `withdraw(amount)`
    * `checkBalance()`
* Use switch for command routing
* Protect against overdrafts

---

### **8. Fibonacci Calculator (Iterative + Recursive)**

🧩 **Concepts**: recursion, loops, `defer`, benchmarking logic
📌 **Task**:

* Write:

  ```go
  func fibRec(n int) int
  func fibIter(n int) int
  ```
* Print and compare their outputs
* Optionally, time the functions and `defer` the print of elapsed time

---

### **9. Math Expression Evaluator**

🧩 **Concepts**: switch, string input, parsing, function dispatch
📌 **Task**:

* Input: `add 3 4` or `mul 2 5`
* Output: `7` or `10`
* Route to correct function using `switch` or map of functions.

---

### **10. Command Line Contact Book (Bonus Project)**

🧩 **Concepts**: CLI args, loop, map, functions, I/O
📌 **Features**:

* Add contact (`add John 12345`)
* Search contact (`search John`)
* List all contacts
* Use loop and `switch` for command dispatch

---
