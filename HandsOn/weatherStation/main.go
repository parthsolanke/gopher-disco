package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TemperatureReading struct {
	Value float64
	Time  time.Time
}

type HumidityReading struct {
	Value float64
	Time  time.Time
}

type WindReading struct {
	Value float64
	Time  time.Time
}

func main() {
	tempChan := make(chan TemperatureReading)
	humidityChan := make(chan HumidityReading)
	windChan := make(chan WindReading)

	go simulateTemperatureSensor(tempChan)
	go simulateHumiditySensor(humidityChan)
	go simulateWindSensor(windChan)

	for {
		select {
		case temp := <-tempChan:
			fmt.Printf("🌡️  Temperature: %.2f°C at %s\n", temp.Value, temp.Time.Format(time.RFC1123))
		case hum := <-humidityChan:
			fmt.Printf("💧 Humidity: %.2f%% at %s\n", hum.Value, hum.Time.Format(time.RFC1123))
		case wind := <-windChan:
			fmt.Printf("💨 Wind Speed: %.2f km/h at %s\n", wind.Value, wind.Time.Format(time.RFC1123))
		}
	}
}

func simulateTemperatureSensor(ch chan<- TemperatureReading) {
	for {
		reading := TemperatureReading{
			Value: 15 + rand.Float64()*10, // 15°C to 25°C
			Time:  time.Now(),
		}
		ch <- reading
		time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond) // 1-3 sec
	}
}

func simulateHumiditySensor(ch chan<- HumidityReading) {
	for {
		reading := HumidityReading{
			Value: 40 + rand.Float64()*30, // 40% to 70%
			Time:  time.Now(),
		}
		ch <- reading
		time.Sleep(time.Duration(rand.Intn(2500)+1000) * time.Millisecond) // 1-3.5 sec
	}
}

func simulateWindSensor(ch chan<- WindReading) {
	for {
		reading := WindReading{
			Value: 5 + rand.Float64()*20, // 5 to 25 km/h
			Time:  time.Now(),
		}
		ch <- reading
		time.Sleep(time.Duration(rand.Intn(3000)+1000) * time.Millisecond) // 1-4 sec
	}
}
