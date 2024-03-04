package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Creates a bcrypt hash of a password.
// Returns the hash as a string and an error if one occurred.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Verifies the password against a bcrypt hash.
// Returns true if the password matches the hash, false otherwise.
func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
