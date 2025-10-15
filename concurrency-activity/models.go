package main

// User represents a single user's info
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"username"`
}
