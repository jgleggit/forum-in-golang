package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func spawnNewInstance() (*os.Process, error) {
	fmt.Println("Spawning a new instance...")

	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return nil, err
	}

	cmd := exec.Command(executable, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start new instance:", err)
		return nil, err
	}

	return cmd.Process, nil
}

func main() {
	fmt.Printf("Running main process (PID: %d)...\n", os.Getpid())

	// Create a channel to wait for the SIGINT signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	var newProcess *os.Process

	for {
		fmt.Println("Waiting for SIGINT signal...")

		// Wait for the SIGINT signal
		<-sigCh

		fmt.Println("Received SIGINT signal. Initiating restart...")

		// Perform cleanup and finalize tasks if needed

		// Terminate the previous new instance
		if newProcess != nil {
			newProcess.Signal(syscall.SIGTERM)
		}

		// Spawn a new instance of the application
		newProcess, _ = spawnNewInstance()
	}
}




/*
package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func spawnNewInstance() {
	fmt.Println("Spawning a new instance...")

	executable, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return
	}

	cmd := exec.Command(executable, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start new instance:", err)
	}
}

func main() {
	fmt.Printf("Running main process (PID: %d)...\n", os.Getpid())

	for {
		fmt.Println("Waiting for SIGINT signal...")

		// Create a channel to wait for the SIGINT signal
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT)

		// Wait for the SIGINT signal
		<-sigCh

		fmt.Println("Received SIGINT signal. Initiating restart...")

		// Perform cleanup and finalize tasks if needed

		// Spawn a new instance of the application
		spawnNewInstance()

		// Terminate the loop by ending the process
		return
	}
}
*/
