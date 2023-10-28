package filelogger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
	"forum-in-golang/logger"
)

type FileLogger struct {
	filePath string
	logger   *logger.Logger
}


func NewFileLogger() (*FileLogger, error) {
	// Check if logs directory exists
	_, err := os.Stat(logger.LogsDirPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(logger.LogsDirPath, 0755)
		if err != nil {
			return nil, err
		}
		fmt.Println("INFO: 'logs' directory created")
	} else if err == nil {
		fmt.Println("INFO: 'logs' directory already exists")
	} else {
		return nil, fmt.Errorf("ERROR: Failed to check 'logs' directory: %v", err)
	}

	now := time.Now()
	fileName := fmt.Sprintf("forum-%04d-%02d-%02d.log", now.Year(), now.Month(), now.Day())
	filePath := filepath.Join(logger.LogsDirPath, fileName)

	// Check if log file already exists
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("INFO: '%s' file already exists\n", fileName)
	} else if os.IsNotExist(err) {
		fmt.Printf("INFO: '%s' file created\n", fileName)

	// Create new log file
	file, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Failed to create log file: %v", err)
	}
	file.Close()
	} else {
		return nil, fmt.Errorf("ERROR: Failed to check forum log file: %v", err)
	}
	return &FileLogger{
		filePath: filePath,
		logger:   logger.NewLogger(),
	}, nil
}


func (f *FileLogger) Log(message string) {
	file, err := os.OpenFile(f.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		f.logger.Error(err)
		return
	}
	defer file.Close()

	fileLogger := log.New(file, "", log.LstdFlags)
	fileLogger.Println(message)
}
