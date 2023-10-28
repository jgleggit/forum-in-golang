package main

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	databaseName := "forum-db.sqlite3"
	databasePath := filepath.Join("databases", databaseName)

	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		fmt.Printf("Failed to open database: %v\n", err)
		log.Fatal(err)
		//db.SetMaxOpenConns(1)
		//sql.Open("file:locked.sqlite?cache=shared", databasePath)
	}
	
	
	//defer db.Close()

	fmt.Println("Database opened successfully!")

	// Check if 'users' table exists
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='users'")
	if err != nil {
		fmt.Printf("Failed to execute SQL query: %v\n", err)
		log.Fatal(err)
	}
	//defer rows.Close()

	if !rows.Next() {
		fmt.Printf("Table 'users' does not exist.\n")
		// 'users' table does not exist, so create it
		createUsersTable := `
			CREATE TABLE IF NOT EXISTS users (
				userID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
				userEmail VARCHAR(100) NOT NULL,
				userName VARCHAR(60) NOT NULL,
				userPassword VARCHAR(60) NOT NULL,
				sessionID CHAR(32) NOT NULL
			);`
			// if using userUUID instead of userID replace above line with this:
			// userUUID CHAR(36) PRIMARY KEY NOT NULL,

		_, err = db.Exec(createUsersTable)
		if err != nil {
			fmt.Printf("Failed to execute SQL statement for creating the 'user' table: %v\n", err)
			log.Fatal(err)
		}
		fmt.Println("Table 'users' created successfully!")
		//rows.Close()
		//db.Close()
	} else {
		fmt.Println("Table 'users' already exists.")
	}


	// Check if 'categories' table exists
	rows2, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='categories'")
	if err != nil {
		fmt.Printf("Failed to execute SQL query: %v\n", err)
		log.Fatal(err)
	}
	//defer rows.Close()
	
	if !rows2.Next() {
		fmt.Printf("Table 'categories' does not exist.\n")
		// 'categories' table does not exist, so create it
		createCategoriesTable := `
			CREATE TABLE IF NOT EXISTS categories (
				categoryID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
				postID INTEGER NOT NULL,
				categoryA VARCHAR(60) NOT NULL,
				categoryB VARCHAR(60) NOT NULL,
				categoryC VARCHAR(60) NOT NULL,
				categoryD VARCHAR(60) NOT NULL,
				categoryE VARCHAR(60) NOT NULL,
				sessionID VARCHAR(32)
			);`

		_, err = db.Exec(createCategoriesTable)
		if err != nil {
			fmt.Printf("Failed to execute SQL statement for creating the 'categories' table: %v\n", err)
			log.Fatal(err)
		}
		fmt.Println("Table 'categories' created successfully!")
		//rows.Close()
		//db.Close()
	} else {
		fmt.Println("Table 'categories' already exists.")
	}



	// Check if 'posts' table exists
	rows, err = db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='posts'")
	if err != nil {
		fmt.Printf("Failed to execute SQL query: %v\n", err)
		log.Fatal(err)
	}
	//defer rows.Close()

	if !rows.Next() {
		fmt.Printf("Table 'posts' does not exist.\n")
		// 'posts' table does not exist, so create it
		createPostsTable := `
			CREATE TABLE IF NOT EXISTS posts (
				postID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
				userID INTEGER NOT NULL,
				postDate VARCHAR(13) NOT NULL,
				titlePost VARCHAR(200) NOT NULL,
				contentPost VARCHAR(1000) NOT NULL,
				categoryID INTEGER NOT NULL,
				sessionID VARCHAR(32)
			);`

		_, err = db.Exec(createPostsTable)
		if err != nil {
			fmt.Printf("Failed to execute SQL statement for creating the 'posts' table: %v\n", err)
			log.Fatal(err)
		}
		fmt.Println("Table 'posts' created successfully!")
		//rows.Close()
		//db.Close()
	} else {
		fmt.Println("Table 'posts' already exists.")
	}

	


		// Check if 'comments' table exists
		rows, err = db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='comments'")
		if err != nil {
			fmt.Printf("Failed to execute SQL query: %v\n", err)
			log.Fatal(err)
		}
		//defer rows.Close()
	
		if !rows.Next() {
			fmt.Printf("Table 'comments' does not exist.\n")
			// 'comments' table does not exist, so create it
			createCommentsTable := `
				CREATE TABLE IF NOT EXISTS comments (
					commentID INTEGER PRIMARY KEY AUTOINCREMENT,
					postID INTEGER NOT NULL,
					userID INTEGER NOT NULL,
					commentDate VARCHAR(10) NOT NULL,
					contentComment VARCHAR(1000) NOT NULL,
					isAuthor VARCHAR(10) NOT NULL,
					categoryID INTEGER NOT NULL,
					sessionID VARCHAR(32)
				);`
	
			_, err = db.Exec(createCommentsTable)
			if err != nil {
				fmt.Printf("Failed to execute SQL statement for creating the 'comments' table: %v\n", err)
				log.Fatal(err)
			}
			fmt.Println("Table 'comments' created successfully!")
			//rows.Close()
		} else {
			fmt.Println("Table 'comments' already exists.")
		}

		// Check if 'postLikes' table exists
		rows, err = db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='postLikes'")
		if err != nil {
			fmt.Printf("Failed to execute SQL query: %v\n", err)
			log.Fatal(err)
			}
			//defer rows.Close()
			
			if !rows.Next() {
				fmt.Printf("Table 'postLikes' does not exist.\n")
				// 'postLikes' table does not exist, so create it
				createpostLikesTable := `
					CREATE TABLE IF NOT EXISTS postLikes (
						postLikesID INTEGER PRIMARY KEY AUTOINCREMENT,
						postID INTEGER NOT NULL,
						userID INTEGER NOT NULL,
						isAuthor VARCHAR(10) NOT NULL,
						sessionID VARCHAR(32)
					);`
			
				_, err = db.Exec(createpostLikesTable)
				if err != nil {
					fmt.Printf("Failed to execute SQL statement for creating the 'postLikes' table: %v\n", err)
					log.Fatal(err)
				}
				fmt.Println("Table 'postLikes' created successfully!")
				//rows.Close()
			} else {
				fmt.Println("Table 'postLikes' already exists.")
			}


		// Check if 'commentLikes' table exists
		rows, err = db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='commentLikes'")
		if err != nil {
			fmt.Printf("Failed to execute SQL query: %v\n", err)
			log.Fatal(err)
		}
		//defer rows.Close()
			
			if !rows.Next() {
				fmt.Printf("Table 'commentLikes' does not exist.\n")
				// 'commentLikes' table does not exist, so create it
				createcommentLikesTable := `
					CREATE TABLE IF NOT EXISTS commentLikes (
						commentLikesID INTEGER PRIMARY KEY AUTOINCREMENT,
						commentID INTEGER NOT NULL,
						postID INTEGER NOT NULL,
						userID INTEGER NOT NULL,
						isAuthor VARCHAR(10) NOT NULL,
						sessionID VARCHAR(32)
					);`

				_, err = db.Exec(createcommentLikesTable)
				if err != nil {
					fmt.Printf("Failed to execute SQL statement for creating the 'commentLikes' table: %v\n", err)
					log.Fatal(err)
				}
				fmt.Println("Table 'commentLikes' created successfully!")
				//rows.Close()
				} else {
					fmt.Println("Table 'commentLikes' already exists.")
				}
				//rows.Close()
	/*

	// Check if 'comments' table exists
	rows, err = db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name='comments'")
	if err != nil {
		fmt.Printf("Failed to execute SQL query: %v\n", err)
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		fmt.Println("Table 'comments' does not exist.")
		// 'comments' table does not exist, so create it
		createCommentsTable := `
			CREATE TABLE IF NOT EXISTS comments (
				commentID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
				postID INTEGER NOT NULL,
				userID INTEGER NOT NULL,
				commentDate VARCHAR(10),
				contentComment VARCHAR(300),
				isAuthor VARCHAR(30)
			);`
			
			param := createCommentsTable
			paramName := fmt.Sprintf("%T" , param)
			_, err = db.Exec(param)
			if err != nil {
				fmt.Printf("Failed to execute SQL statement triggered by %q: %v\n%v\n", paramName, param, err)
				log.Fatal(err)
			}
			
			
			
		fmt.Println("Table 'comments' created successfully!")
		rows.Close()
		//db.Exec( )
	} else {
		fmt.Println("Table 'comments' already exists.")
	}
*/
	fmt.Println("Closing the database.")
}
