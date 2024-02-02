package utils

import (
	"fmt"
	"log"
)

func LogError(err error) {
	if err != nil {
		log.Printf("[ERROR] %v", err)
	}
}

func LogInfo(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Printf("[INFO] %s", message)
}

func LogFatal(format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	log.Fatalf("[FATAL] %s", message)
}
