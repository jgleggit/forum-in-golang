package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	//"runtime/debug"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	databaseName := "forum-db.sqlite3"
	databasePath := filepath.Join("databases", databaseName)
	directoryPath := filepath.Dir(databasePath)

	fmt.Printf("Checking if directory '%s' exists\n", directoryPath)

	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' does not exist. Creating...\n", directoryPath)

		// Create the directory
		err = os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			fmt.Printf("Failed to create directory '%s': %v\n", directoryPath, err)
			log.Fatal(err)
		}
	} else if err == nil {
		fmt.Printf("Directory '%s' already exists.\n", directoryPath)
	} else {
		log.Fatal(err)
	}

	fmt.Printf("Checking if SQLite database file '%s' exists\n", databaseName)

	// Check if the database file exists
	_, err = os.Stat(databasePath)
	if os.IsNotExist(err) {
		fmt.Printf("Database file '%s' does not exist. Creating...\n", databaseName)

		// Create the SQLite database file
		file, err := os.Create(databasePath)
		if err != nil {
			fmt.Printf("Failed to create database file '%s': %v\n", databaseName, err)
			log.Fatal(err)
		}
		file.Close()

		fmt.Printf("SQLite database '%s' created successfully!\n", databaseName)
	} else if err != nil {
		fmt.Printf("Error checking database file '%s': %v\n", databaseName, err)
		log.Fatal(err)
	} else {
		fmt.Printf("SQLite database '%s' already exists.\n", databaseName)
	}

	fmt.Printf("Attempting to open and connect to SQLite database '%s'\n", databaseName)

	// Open and connect to the database
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		fmt.Printf("Failed to open and connect to SQLite database '%s': %v\n", databaseName, err)
	} else {
		// defer db.Close()

		// 	//_ = sql.Open("sqlite3", databasePath)
		// 	bi, ok := debug.ReadBuildInfo()
		// 	if !ok {
		// 		log.Printf("Failed to read build info")
		// 		return
		// 	}
		// 	for _, dep := range bi.Deps {
		// 		fmt.Printf("Dep: %+v\n", dep)
		// }
		fmt.Printf("Opened and connected to SQLite database '%s' successfully!\n", databaseName)

		//db, err := sql("sqlite3", databasePath)
		//db.Close()

		
		// Perform additional operations with the database as needed
	}
	db.Close()
	fmt.Printf("Closed Sqlite database '%s' successfully.\n", databaseName)
}
