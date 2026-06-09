// This example demonstrates the intended integration style between
// helper/v2 and github.com/akula410/connect.
//
// Adjust connect API calls to the actual github.com/akula410/connect major version.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/akula410/helper/v2/convert"
	"github.com/akula410/helper/v2/password"
	"github.com/akula410/helper/v2/random"
)

// connectDB is a placeholder for the actual connect.Open call.
// Replace with: db, err := connect.Open(connect.Config{...})
func connectDB() error {
	// This example demonstrates the intended integration style.
	// Adjust connect API calls to the actual github.com/akula410/connect major version.
	return nil
}

// UserRow simulates a row returned from the database.
type UserRow struct {
	ID           any
	Email        string
	PasswordHash string
}

func main() {
	ctx := context.Background()
	_ = ctx

	if err := connectDB(); err != nil {
		log.Fatal(err)
	}

	// --- Simulate fetching a user from the DB ---
	// In practice, use connect + builder to query the DB:
	//   row := db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = ?", email)
	userRow := UserRow{
		ID:           "42",
		Email:        "alice@example.com",
		PasswordHash: "",
	}

	// --- Hash a password before storing ---
	hash, err := password.Hash("super_secret")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Password hash:", hash)

	// --- Verify password on login ---
	ok := password.Verify(userRow.PasswordHash, "super_secret")
	fmt.Println("Login ok:", ok)

	// --- Convert DB values safely ---
	userID := convert.Int64(userRow.ID, 0)
	fmt.Println("User ID:", userID)

	// --- Generate a request correlation ID ---
	requestID, err := random.UUIDv4()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Request ID:", requestID)

	// --- Generate a session token ---
	token, err := random.Token(32)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Session token:", token)
}
