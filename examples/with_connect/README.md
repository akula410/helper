# Integration example: helper/v2 + akula410/connect

This directory shows the **intended integration pattern** between `helper/v2` and
`github.com/akula410/connect`. Because the connect package API may differ between
major versions, this example is provided as documentation rather than a
compile-ready program. Adjust the connect calls to the actual API of the major
version you are using.

## Scenario

- Connect to MySQL via `github.com/akula410/connect`.
- Fetch a user row.
- Verify the password using `helper/v2/password`.
- Convert DB values with `helper/v2/convert`.
- Generate a session token with `helper/v2/random`.

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    // Adjust connect import to the actual major version.
    // connect "github.com/akula410/connect/v2"

    "github.com/akula410/helper/v2/convert"
    "github.com/akula410/helper/v2/password"
    "github.com/akula410/helper/v2/random"
)

func main() {
    ctx := context.Background()

    // --- Open DB connection ---
    // db, err := connect.Open(connect.Config{
    //     Host:     "127.0.0.1",
    //     Port:     3306,
    //     User:     "root",
    //     Password: "secret",
    //     Database: "myapp",
    // })
    // if err != nil { log.Fatal(err) }
    // defer db.Close()

    // --- Hash a password before storing ---
    hash, err := password.Hash("super_secret")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Password hash:", hash)

    // --- Fetch user row (adjust to actual connect/builder query API) ---
    // row := db.QueryRowContext(ctx, "SELECT id, email, password FROM users WHERE email = ?", "alice@example.com")
    // var userID any
    // var passwordHash string
    // if err := row.Scan(&userID, nil, &passwordHash); err != nil { log.Fatal(err) }

    // Simulate fetched values:
    userID := "42"
    passwordHash := hash

    // --- Verify password on login ---
    ok := password.Verify(passwordHash, "super_secret")
    fmt.Println("Login ok:", ok)

    // --- Convert DB values safely ---
    id := convert.Int64(userID, 0)
    fmt.Println("User ID:", id)

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

    _ = ctx
}
```
