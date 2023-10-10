func RegisterUser(username, useremail, userpass string) {
	// Define the SQL query for user registration
	var registrationQuery = "INSERT INTO users(username, useremail, userpass) VALUES(?, ?, ?)"

	// Generate a hash of the user's password with a cost factor of 12
	var hashedPassword []byte
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userpass), 12)

	// Prepare the SQL statement for registration
	registrationStatement, err := Database.Prepare(registrationQuery)
	HandleError(err)

	defer registrationStatement.Close()

	// Execute the registration statement with user information
	registrationStatement.Exec(username, useremail, hashedPassword)
}
