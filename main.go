package repos

import "log"

func LogInfo(message string) {
	log.Printf("INFO - %v", message)
}

func SayHello() string {
	return "Hello World"
}
