package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

const sharedPIDFile = "shared-pid.txt"

func writeSharedPID(pid int) error {
	pidStr := strconv.Itoa(pid)
	return ioutil.WriteFile(sharedPIDFile, []byte(pidStr), 0644)
}

func readSharedPID() (int, error) {
	pidBytes, err := ioutil.ReadFile(sharedPIDFile)
	if err != nil {
		return 0, err
	}

	pidStr := strings.TrimSpace(string(pidBytes))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func spawnNewInstance(executable string) (*os.Process, error) {
	fmt.Println("Spawning a new instance of file 1...")

	cmd := exec.Command(executable)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Failed to start new instance:", err)
		return nil, err
	}

	return cmd.Process, nil
}

func forwardSIGINT(pid int) {
	fmt.Printf("Forwarding SIGINT to PID %d...\n", pid)
	if err := syscall.Kill(pid, syscall.SIGINT); err != nil {
		fmt.Printf("Failed to forward SIGINT to PID %d: %v\n", pid, err)
	}
}

func main() {
	pid := os.Getpid()
	fmt.Printf("running file2.go with PID[%d] ...\n", pid)

	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
		return
	}

	executable := executablePath

	// Create a channel to wait for the SIGINT signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	// Write the initial shared PID to the file
	if err := writeSharedPID(pid); err != nil {
		fmt.Println("Failed to write shared PID:", err)
		return
	}

	var newProcess *os.Process

	for {
		sharedPID, err := readSharedPID()
		if err != nil {
			fmt.Println("Failed to read shared PID:", err)
		} else if sharedPID != 0 {
			fmt.Printf("file1.go is running with PID[%d]\n", sharedPID)
			fmt.Printf("Forwarding SIGINT to file1.go (PID: %d)...\n", sharedPID)
			forwardSIGINT(sharedPID)
		}

		<-sigCh

		fmt.Printf("Received SIGINT signal from file1.go with PID[%d]). Initiating restart...\n", sharedPID)

		if newProcess != nil {
			fmt.Printf("Terminating previous instance of file1.go (PID: %d)...\n", newProcess.Pid)
			newProcess.Signal(syscall.SIGTERM)
			newProcess.Wait()
		}

		newProcess, _ = spawnNewInstance(executable)
		if err := writeSharedPID(newProcess.Pid); err != nil {
			fmt.Println("Failed to write shared PID:", err)
		}

		fmt.Printf("file1.go is running with PID[%d]\n", newProcess.Pid)
		fmt.Printf("Forwarding SIGINT to file1.go (PID: %d)...\n", newProcess.Pid)
		forwardSIGINT(newProcess.Pid)
	}
}
