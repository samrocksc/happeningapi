package main

// TransformedUsers - contains the expected output for a user if they exist.
type TransformedUsers struct {
	ID          int `json:"id"`
	IsValidated bool
	Version     int
	Email       string
	Username    string
}
