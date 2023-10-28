package main

import (
	"fmt"
	"os/exec"
)

func main() {
	runOSCommand()
}

func runOSCommand() {
	// Start the process in the background
	cmd := exec.Command("sleep", "10")
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Process started with PID: %d\n", cmd.Process.Pid)

	// Wait for the process to complete
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Process completed.")
}