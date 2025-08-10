package main

import (
	"fmt"
	"time"
)

// The select statement lets you wait on multiple channel operations simultaneously. It's like a switch statement, but for channels. Whichever channel is ready first gets processed.
func fireStation(emergencyChannel chan string) {
	time.Sleep(2 * time.Second)
	emergencyChannel <- "🔥 House fire on Oak Street!"
}

func policeStation(emergencyChannel chan string) {
	time.Sleep(1 * time.Second)
	emergencyChannel <- "🚔 Car accident on Main Street!"
}

func hospital(emergencyChannel chan string) {
	time.Sleep(3 * time.Second)
	emergencyChannel <- "🏥 Medical emergency on Pine Avenue!"
}

func main() {
	fireEmergency := make(chan string)
	policeEmergency := make(chan string)
	medicalEmergency := make(chan string)

	// All emergency services are on standby
	go fireStation(fireEmergency)
	go policeStation(policeEmergency)
	go hospital(medicalEmergency)

	fmt.Println("📞 Emergency Dispatcher waiting for calls...")

	// Handle the first emergency that comes in
	select {
	case fireCall := <-fireEmergency:
		fmt.Printf("📞 Dispatcher: URGENT! %s\n", fireCall)
	case policeCall := <-policeEmergency:
		fmt.Printf("📞 Dispatcher: URGENT! %s\n", policeCall)
	case medicalCall := <-medicalEmergency:
		fmt.Printf("📞 Dispatcher: URGENT! %s\n", medicalCall)
	}

	fmt.Println("📞 Emergency response dispatched!")
}
