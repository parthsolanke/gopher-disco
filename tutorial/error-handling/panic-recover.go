package main

import "fmt"

// only use panic and recover when necessary and critical
// Something so wrong happened that continuing is impossible
//- Array index out of bounds (programming error)
//- Nil pointer dereference (programming error)
//- Configuration so broken the app can't start

// Emergency response team 🚑
func rescue() {
	// recover can also be used when not explicit using panic
	if r := recover(); r != nil {
		fmt.Println("🚑 Caught panic:", r)
		fmt.Println("✅ Program continues safely!")
	}
}

// Function that might cause trouble 🔥
func risky(data []int, index int) {
	defer rescue() // Always have rescue ready
	if index >= len(data) {
		panic("Index out of bounds! 🔥") // Something catastrophic
	}
	fmt.Println("Data:", data[index])
}

func main() {
	fmt.Println("🎬 Starting risky operation...")
	risky([]int{1, 2, 3}, 10) // Bad index -> panic -> recover
	fmt.Println("🎉 We survived the panic!")
}
