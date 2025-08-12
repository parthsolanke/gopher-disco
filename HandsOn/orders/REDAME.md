Concurrent Order Processing System
Background
You are tasked with designing a simplified concurrent order processing service for a small e-commerce platform.
The system will accept order details from the user, process multiple orders in parallel, and handle real-world failures like payment issues and stock shortages.
Example output -
Enter number of orders: 3

--- Order #1 ---
Order ID: A101
Items (comma separated): Phone, Charger
Amount: 250.50

--- Order #2 ---
Order ID: A102
Items (comma separated): Laptop, Mouse
Amount: 1300

--- Order #3 ---
Order ID: A103
Items (comma separated): Tablet
Amount: 499.99

[Worker-1] Processing Order#A101 ... ✅ Payment Success
[Worker-2] Processing Order#A102 ... ❌ Out of Stock
[Worker-1] Processing Order#A103 ... ❌ Payment Failed

--- Errors ---
Error: payment failed for order A103: gateway timeout
Error: item Laptop out of stock for order A102
Constraints
1. Payment failures occur randomly (~30% chance).
2. Stock levels are pre-defined in a map.
3. Only one goroutine may modify stock at a time.
4. All goroutines should finish before the program exits.









