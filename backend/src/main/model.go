package main

// UserProfile is the expected data for a post request on the /signup route
type UserProfile struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// SessionDB just simulates a database
type SessionDB struct {
	sessionID string
	userID    string
}
