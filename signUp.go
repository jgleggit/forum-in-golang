func RegisterUser(username, useremail, userpass string) {
	// Define the SQL query for user registration
	var registerUserQuery = "INSERT INTO users(username, useremail, userpass) VALUES(?, ?, ?)"

	// Generate a hash of the user's password with a cost factor of 12
	var hashedPassword []byte
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userpass), 12) // Set cost to 12

	// Prepare the SQL statement for user registration
	registerUserStatement, err := Database.Prepare(registerUserQuery)
	HandleError(err)

	defer registerUserStatement.Close()

	// Execute the user registration statement with user information
	registerUserStatement.Exec(username, useremail, hashedPassword)
}
