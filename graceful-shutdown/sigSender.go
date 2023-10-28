package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Get the process ID of the receiver program
	receiverPID := 12345 // Replace with the actual process ID of the receiver

	// Send a signal to the receiver
	if err := syscall.Kill(receiverPID, syscall.SIGUSR1); err != nil {
		fmt.Println("Failed to send signal:", err)
	}

	fmt.Println("Restart signal sent to receiver")
}
