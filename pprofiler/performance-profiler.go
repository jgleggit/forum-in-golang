package main

import (
	_ "net/http/pprof"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	ErrorLevel
)

var logPrefixes = [...]string{
	"[INFO] ",
	"[ERROR] ",
}

func logPrintf(level LogLevel, format string, v ...interface{}) {
	fmt.Printf(logPrefixes[level]+format+"\n", v...)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you have requested: %s\n", r.URL.Path)
}


func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

	// Create a channel to signal the web server to shut down
	serverDoneCh := make(chan struct{})

	// Create a channel to signal the whole app to shut down
	appDoneCh := make(chan struct{})

	// Create an HTTP server
	server := &http.Server{Addr: "localhost:8080"}

	// Start the HTTP server in a goroutine
	go func() {
		logPrintf(InfoLevel, "HTTP server started on port 8080 ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logPrintf(ErrorLevel, "HTTP server failed to start: %v ...", err)
		}
		close(serverDoneCh)
	}()

	// Create a channel to wait for termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT) // Add SIGQUIT

	// Print the URL for pprof tools
	logPrintf(InfoLevel, "Access the pprof tools by opening\n       http://localhost:8080/debug/pprof/ ...")

	// Wait for a signal
	sig := <-sigCh
	logPrintf(InfoLevel, "Received %s signal. Shutting down ...", sig)

	// Signal the web server to shut down
	if err := server.Shutdown(nil); err != nil {
		logPrintf(ErrorLevel, "Error while shutting down server: %v ...", err)
	} else {
		logPrintf(InfoLevel, "HTTP server received signal to shut down ...")
	}

	// Wait for the server to finish
	<-serverDoneCh

	// Printf message indicating that the HTTP server has shutdown gracefully
	logPrintf(InfoLevel, "HTTP server shutdown gracefully ...")

	// Close the appDoneCh to signal that the application is shutting down
	close(appDoneCh)

	// Wait for the appDoneCh to be closed
	<-appDoneCh

	logPrintf(InfoLevel, "Application shut down gracefully...")
}

/*
package main

import (
	_ "net/http/pprof"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

type LogLevel int

const (
	InfoLevel LogLevel = iota
	ErrorLevel
)

var logPrefixes = [...]string{
	"[INFO] ",
	"[ERROR] ",
}

func logPrintf(level LogLevel, format string, v ...interface{}) {
	fmt.Printf(logPrefixes[level]+format+"\n", v...)
}

func main() {
	// Create a channel to signal the web server to shut down
	serverDoneCh := make(chan struct{})

	// Create a channel to signal the whole app to shut down
	appDoneCh := make(chan struct{})

	// Create an HTTP server
	server := &http.Server{Addr: "localhost:8080"}

	// Start the HTTP server in a goroutine
	go func() {
		logPrintf(InfoLevel, "HTTP server started on port 8080 ...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logPrintf(ErrorLevel, "HTTP server failed to start: %v ...", err)
		}
		close(serverDoneCh)
	}()

	// Create a channel to wait for termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Print the URL for pprof tools
	logPrintf(InfoLevel, "Access the pprof tools by opening\n       http://localhost:8080/debug/pprof/ ...")

	// Wait for a signal
	sig := <-sigCh
	logPrintf(InfoLevel, "Received %s signal. Shutting down ...", sig)

	// Signal the web server to shut down
	if err := server.Shutdown(nil); err != nil {
		logPrintf(ErrorLevel, "Error while shutting down server: %v ...", err)
	} else {
		logPrintf(InfoLevel, "HTTP server received signal to shut down ...")
	}

	// Wait for the server to finish
	<-serverDoneCh

	// Printf message indicating that the HTTP server has shutdown gracefully
	logPrintf(InfoLevel, "HTTP server shutdown gracefully ...")

	// Close the appDoneCh to signal that the application is shutting down
	close(appDoneCh)

	// Wait for the appDoneCh to be closed
	<-appDoneCh

	logPrintf(InfoLevel, "Application shut down gracefully...")
}
*/


/*
package main

import (
	_ "net/http/pprof"
	"net/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Access the pprof tools by opening
	// http://localhost:8080/debug/pprof/

	// Create a channel to signal the web server to shut down
	serverDoneCh := make(chan struct{})

	// Create a channel to signal the whole app to shut down
	appDoneCh := make(chan struct{})

	// Create an HTTP server
	server := &http.Server{Addr: "localhost:8080"}

	// Start the HTTP server in a goroutine
	go func() {
		log.Println("Server started on port 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("HTTP server failed to start: %v", err)
		}
		close(serverDoneCh)
	}()

	// Create a channel to wait for termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for a signal
	sig := <-sigCh
	log.Printf("Received %s signal. Shutting down gracefully...", sig)

	// Signal the web server to shut down
	if err := server.Shutdown(nil); err != nil {
		log.Printf("Error while shutting down server: %v", err)
	} else {
		log.Println("HTTP server received signal to shut down")
	}

	// Wait for the server to finish
	<-serverDoneCh

	// Printf message indicating that the HTTP server has shutdown gracefully
	log.Println("HTTP server shutdown gracefully")

	// Close the appDoneCh to signal that the application is shutting down
	close(appDoneCh)

	// Wait for the appDoneCh to be closed
	<-appDoneCh

	log.Println("Application gracefully shut down")
}
*/