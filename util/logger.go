package util

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "[CFB-API] ", log.LstdFlags|log.Lshortfile)
}

func LogDBError(operation string, err error) {
	if err != nil {
		Logger.Printf("DB Error in %s: %v", operation, err)
	}
}
