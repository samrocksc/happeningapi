package main

import (
	"golang.org/x/crypto/bcrypt"
)

// ComparePassword - Compares password and returns a boolean
func ComparePassword(a, b string) bool {
	hashedPassword := []byte(a)
	password := []byte(b)
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password).Error; err != nil {
		return false
	}
	return true
}
