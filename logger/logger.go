package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Logger exported
var Logger *log.Logger

func init() {
	absPath, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/goshard-log.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	Logger = log.New(generalLog, "Goshard Logger:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
