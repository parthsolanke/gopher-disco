package logger

import "fmt"

func init() {
	fmt.Println("Logger package initialized")
}

func Log(msg string) {
	fmt.Println("LOG: " + msg)
}
