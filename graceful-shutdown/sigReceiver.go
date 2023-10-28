package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to wait for the signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGUSR1)

	// Wait for the signal
	sig := <-signalCh

	// Print confirmation message
	fmt.Println("Received restart signal (SIGUSR1)")

	// Perform restart logic here
	// For this example, we'll just simulate a restart by exiting
	fmt.Println("Performing restart...")
	os.Exit(0)
}
