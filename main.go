// main.go
package main

import (
	"fmt"

	"forum-in-golang/filelogger"
	applogger "forum-in-golang/logger"
)

var appLog *applogger.Logger

func main() {
	runApplication()
	compareInts(10, 11)
	checkAge(5)
}

func runApplication() {
	appLog = applogger.NewLogger()

	fileLogger, err := filelogger.NewFileLogger()
	if err != nil {
		appLog.Error(err)
		return
	}

	appLog.Info("Custom Info log message")
	appLog.Warning("Custom Warning log message")
	appLog.Error("Custom Error log message")

	appLog.Fatal("Custom Fatal log message")

	fileLogger.Log("Log message to file")

	appLog.Fatal(fmt.Sprintf("FATAL: %s", "No further details"))
}

// compareInts()
func compareInts(x int, y int) {
	if x == y {
		appLog.Info("Integers are equal")
	} else {
		appLog.Warning("Integers are not equal")
	}
}

// checkAge()
func checkAge(n int) {
	if n < 18 {
		appLog.Warning(fmt.Sprintf("WARNING: given age '%d' is under 18!", n))
	} else {
		appLog.Info(fmt.Sprintf("INFO: given age '%d' is 18 or over.", n))
	}
}
