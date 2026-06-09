// Package password provides bcrypt-based password hashing and verification.
package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// ErrEmptyPassword is returned when an empty password is provided.
var ErrEmptyPassword = errors.New("password: empty password")

// Hash generates a bcrypt hash of the given plaintext password using bcrypt.DefaultCost.
// Returns (hash, nil) on success or ("", error) on failure.
func Hash(plain string) (string, error) {
	return HashWithCost(plain, bcrypt.DefaultCost)
}

// HashWithCost generates a bcrypt hash with a custom cost factor.
// cost must be between bcrypt.MinCost and bcrypt.MaxCost.
func HashWithCost(plain string, cost int) (string, error) {
	if plain == "" {
		return "", ErrEmptyPassword
	}
	b, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Verify reports whether plain matches the bcrypt hash.
// Returns false on any error (invalid hash, wrong password, etc.) — never panics.
func Verify(hash string, plain string) bool {
	if hash == "" || plain == "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}
