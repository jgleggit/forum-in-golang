package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// Define a struct that simulates a mock database
type MockDB struct {
    Users map[string]User // Simulate a table of users
}

// User represents a user in the mock database
type User struct {
    Username string
    Useremail string
    Userpass []byte
}

// Create a new instance of the mock database
var mockDB = &MockDB{
    Users: make(map[string]User),
}

// Replace the real database with the mock database
var Database = mockDB

// RegisterUser function (from the original code)
func RegisterUser(username, useremail, userpass string) error {
    // Generate a hash of the user's password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userpass), 12) // Set cost to 12
    if err != nil {
        return err
    }

    // Store the user data in the mock database
    mockDB.Users[username] = User{
        Username: username,
        Useremail: useremail,
        Userpass: hashedPassword,
    }

    return nil
}

func TestRegisterUser_HappyPath(t *testing.T) {
    // Define test input
    username := "testuser"
    useremail := "test@example.com"
    userpass := "testpassword"

    // Call the RegisterUser function
    err := RegisterUser(username, useremail, userpass)

    // Perform assertions using the mock database
    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    // Verify that the user was successfully registered in the mock database
    registeredUser, exists := mockDB.Users[username]
    if !exists {
        t.Errorf("Expected user to be registered, but it wasn't.")
    }

    // Verify other details if needed
    if registeredUser.Username != username {
        t.Errorf("Expected username to be %s, but got %s", username, registeredUser.Username)
    }
}

func TestRegisterUser_ErrorHandling(t *testing.T) {
    // Introduce an error condition (e.g., mock a hashing error)
    // For this example, let's simulate a hashing error by passing an empty password

    // Define test input
    username := "testuser"
    useremail := "test@example.com"
    userpass := "" // Empty password intentionally

    // Call the RegisterUser function
    err := RegisterUser(username, useremail, userpass)

    // Perform assertions to verify that error handling works as expected
    if err == nil {
        t.Error("Expected an error, but got nil.")
    } else if err.Error() != "crypto/bcrypt: hashedSecret too short to be a bcrypted password" {
        t.Errorf("Expected a specific error message, but got: %v", err)
    }
}

func main() {
    // Run the tests
    testing.Main(nil, nil, nil, nil)
}
