package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Define a struct for custom errors with informative messages
type UserRegistrationError struct {
	Message string
}

// Error is a method that defines the behavior of the UserRegistrationError 
// type to satisfy the error interface. It returns the custom error message 
// associated with UserRegistrationError.
//
// Satisfying the error interface means implementing the necessary method to
// make UserRegistrationError compatible with Go's error handling mechanism.
func (e *UserRegistrationError) Error() string {
	return e.Message
}


// Define constants for logging levels
// To-do: Replace with "log/slog" structured logging
const (
	logLevelError = "ERROR: "
	logLevelInfo  = "INFO: "
)

func initLogging() (*os.File, *log.Logger, error) {
	// Create a log file for errors and info
	logFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.APPEND, 0666)
	if err != nil {
		return nil, nil, fmt.Errorf("Error creating log file: %v", err)
	}

	// Create a logger that logs errors to both the file and the terminal
	logger := log.New(io.MultiWriter(logFile, os.Stdout), logLevelError, log.Ldate|log.Ltime|log.Lshortfile)

	return logFile, logger, nil
}

func HandleError(err error, logger *log.Logger, logLevel, errorMessage string) bool {
	if err != nil {
		logger.Printf("%s%s: %v", logLevel, errorMessage, err)
		return true
	}
	return false
}

func RegisterUser(username, useremail, userpass string, logger *log.Logger) error {
	// Define the SQL query for user registration
	var registerUserQuery = "INSERT INTO users(username, useremail, userpass) VALUES(?, ?, ?)"

	
	// CostFactor defines the number of cost factor rounds for password hashing.
	const ( 
    	CostFactor = 12
	)

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(userpass), CostFactor)

	// Check for errors during password hashing
	if hashErr != nil {
    	return fmt.Errorf("Error hashing password: %w", hashErr)
	}

	// Log an "INFO" message for successful password hashing
	logger.Printf("%sPassword hashing successful\n", logLevelInfo)


	// Prepare the SQL statement for user registration
	registerUserStatement, prepareErr := Database.Prepare(registerUserQuery)

	// Check for errors during SQL statement preparation
	if prepareErr != nil {
		return fmt.Errorf("Error preparing SQL statement: %w", prepareErr)
	}
	defer registerUserStatement.Close()

	// Execute the user registration statement with user information
	executionResult, executionError := registerUserStatement.Exec(username, useremail, hashedPassword)
	if executionError != nil {
		return fmt.Errorf("Error executing SQL statement: %w", executionError)
	}

// Log an "INFO" message for successful execution
logger.Printf("%sRegistration SQL statement executed successfully\n", logLevelInfo)


	// Check for errors during SQL statement execution
	if executionError != nil {
		return fmt.Errorf("Error executing SQL statement: %w", executionError)
	}

	// Log an "INFO" message for successful registration
	logger.Printf("%sRegistration successful for user: %s\n", logLevelInfo, username)

	return nil // No error
}

func main() {
	// Initialize logging
	logFile, logger, initErr := initLogging()
	if initErr != nil {
		fmt.Printf("Error initializing logging: %v\n", initErr)
		return
	}
	defer logFile.Close()
	defer logger.Writer().(*ChatGPTos.File).Close()

	// Example usage of RegisterUser function
	registrationError := RegisterUser("testuser", "test@example.com", "testpassword", logger)
	if registrationError != nil {
		logger.Printf("Registration failed: %v\n", registrationError)
	}
}
