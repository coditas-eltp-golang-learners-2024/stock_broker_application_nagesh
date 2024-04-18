package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(logFilePath string) {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}

	Logger = log.New(file, "Stock Broker Application", log.Ldate|log.Ltime)
}
