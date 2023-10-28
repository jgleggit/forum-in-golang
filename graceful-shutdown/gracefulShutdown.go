/*
Simple steps to implement Graceful Shutdown

    Create a signal channel: This channel will be used to receive signals from the operating system when the program is supposed to shutdown.
    Register signal handler: The signal channel should be registered with the signal package to receive signals such as SIGINT and SIGTERM.
    Wait for signal: In the main function of the program, wait for a signal to be received from the signal channel.
    Cleanup: Once a signal is received, perform any necessary cleanup operations before exiting the program.
    Exit: Exit the program in an orderly manner.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a signal channel
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start a goroutine that does some work
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exiting program")
				return
			default:
				fmt.Println("Doing work")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// Wait for a signal to be received
	sig := <-sigCh
	fmt.Println("Received signal:", sig)

	// Cleanup
	fmt.Println("Performing cleanup...")
	time.Sleep(2 * time.Second)

	// Exit
	fmt.Println("Exiting program")
}