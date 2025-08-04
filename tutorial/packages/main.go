package main

// only using the init function of teh logger package
//import _ "packages/logger"
//
//func main() {
//}

// importing and suing logger
import "packages/logger"

func main() {
	logger.Log("hello world")
}
