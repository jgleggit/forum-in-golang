// logger/logger.go
package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	LogsDirPath = "logs"
)

type Logger struct {
	log *log.Logger
}

var appLogger *Logger

func NewLogger() *Logger {
	filePath := getLogFilePath()
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	appLogger = &Logger{
		log: log.New(file, "", log.Ldate|log.Ltime),
	}

	return appLogger
}

func getLogFilePath() string {
	now := time.Now()
	fileName := fmt.Sprintf("forum-%s.log", now.Format("2006-01-02"))
	filePath := filepath.Join(LogsDirPath, fileName)
	return filePath
}

func (l *Logger) Info(message string) {
	l.log.Println("INFO:", message)
}

func (l *Logger) Warning(message string) {
	l.log.Println("WARNING:", message)
}

func (l *Logger) Error(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	errMsg := fmt.Sprint(args...)
	l.log.Println("ERROR:", errMsg)
}

func (l *Logger) Fatal(args ...interface{}) {
	if len(args) == 0 {
		l.log.Println("FATAL: No further details")
		os.Exit(1)
	}

	fatalMsg := fmt.Sprint(args...)
	l.log.Println("FATAL:", fatalMsg)
	os.Exit(1)
}
