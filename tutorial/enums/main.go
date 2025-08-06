package main

import "fmt"

// type
type ServerState int

// enum constant heere you can also initialize strings instead of ints
// in go for performance optimization, this approach is followed to implement enums
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// for small enums do:
// type ServerState string
//const (
//	StateIdle      ServerState = "idle"
//	StateConnected ServerState = "connected"
//	StateError     ServerState = "error"
//	StateRetrying  ServerState = "retrying"
//)

// map to get string value for each state
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// method String implements the fmt.Stringer interface
// values of ServerState can be printed out or converted to strings without explicitly accessing them
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition function making use of the enums
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:

		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
