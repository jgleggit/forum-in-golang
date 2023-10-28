package main

//build-forum-app.go

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	buildName := "forum-app"
	buildPath := filepath.Join("builds", buildName)
	directoryPath := filepath.Dir(buildPath)

	fmt.Printf("Checking if directory '%s' exists ...\n", directoryPath)

	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' does not exist. Creating ...\n", directoryPath)

		// Create the directory
		err = os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory '%s': %v ...\n", directoryPath, err)
			log.Fatal(err)
		}
	} else if err == nil {
		fmt.Printf("Directory '%s' already exists ...\n", directoryPath)
	} else {
		log.Fatal(err)
	}

	fmt.Printf("Checking if Golang app executable file '%s' exists ...\n", buildName)

	// Check if the build file exists
	_, err = os.Stat(buildPath)
	if os.IsNotExist(err) {
		fmt.Printf("Build file '%s' does not exist. Creating ...\n", buildName)

		// Create the Golang app execuatable file
		file, err := os.Create(buildPath)
		if err != nil {
			fmt.Printf("Failed to create build file '%s': %v ...\n", buildName, err)
			log.Fatal(err)
		}
		file.Close()

		fmt.Printf("Golang app executable file '%s' created successfully! ...\n", buildName)
	} else if err != nil {
		fmt.Printf("Error checking executable file '%s': %v ...\n", buildName, err)
		log.Fatal(err)
	} else {
		fmt.Printf("Golang app executable file '%s' already exists ...\n", buildName)
	}
	buildForumApp()
}


func buildForumApp() {
	scriptPath := "./scripts/build-forum-app.sh"

	// Get the directory of the executable
	executableDir, err := os.Executable()
	if err != nil {
		fmt.Printf("Error getting executable directory: %v ...\n", err)
		return
	}
	logFilePath := filepath.Join(filepath.Dir(executableDir), "build-errors.log ...")

	logFile, err := os.Create(logFilePath)
	if err != nil {
		fmt.Printf("Error creating log file: %v ...\n", err)
		return
	}
	defer logFile.Close()

	// Create a logger that writes to both terminal and the log file
	logger := log.New(io.MultiWriter(os.Stdout, logFile), "", log.LstdFlags)

	cmd := exec.Command(scriptPath)

	// Set the logger's output to log level ERROR and above
	logger.SetOutput(logFile)
	logger.SetPrefix("ERROR: ")

	// Print an INFO level message to indicate the start of the build process
	fmt.Printf("INFO: Starting the build process ...\n")
	logger.Printf("INFO: Starting the build process ...\n")

	if startBuild := cmd.Start(); startBuild != nil {
		// Log an ERROR level message for the error
		logger.Printf("ERROR: Error starting process: %v ...\n", startBuild)
		fmt.Printf("ERROR: Error starting process: %v ...\n", startBuild)
	}

	if waitBuild := cmd.Wait(); waitBuild != nil {
		// Log an ERROR level message for the error
		logger.Printf("ERROR: Error waiting for process: %v ...\n", waitBuild)
		fmt.Printf("ERROR: Error waiting for process: %v ...\n", waitBuild)
	}

	// Log an INFO level message indicating the completion of the build
	logger.Println("INFO: Forum app build completed ...")
	fmt.Println("INFO: Forum app build completed ...")
}
