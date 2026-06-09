# Integration example: helper/v2 + akula410/migrations

This directory shows the **intended integration pattern** between `helper/v2` and
`github.com/akula410/migrations`. Because the migrations package API may differ
between major versions, this example is provided as documentation rather than a
compile-ready program. Adjust the migrations calls to the actual API of the major
version you are using.

## Scenario

- Generate a migration-run ID for logging/auditing with `helper/v2/random`.
- Normalise a migration name with `helper/v2/stringsx`.
- Convert env/config values with `helper/v2/convert`.
- Run migrations with `github.com/akula410/migrations`.

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"
    "strings"

    // Adjust migrations import to the actual major version.
    // "github.com/akula410/migrations/v2"

    "github.com/akula410/helper/v2/convert"
    "github.com/akula410/helper/v2/random"
    "github.com/akula410/helper/v2/stringsx"
)

func main() {
    ctx := context.Background()

    // --- Convert env/config values safely ---
    rawVersion := "  1  "
    rawEnv     := "STAGING"

    version := convert.Int(rawVersion, 0)
    env     := strings.ToLower(strings.TrimSpace(rawEnv))
    fmt.Printf("Running migrations for env=%s version=%d\n", env, version)

    // --- Generate a migration-run ID for audit logging ---
    runID, err := random.UUIDv4()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Migration run ID:", runID)

    // --- Normalise a migration name ---
    rawName := "  Add_Users_Table  "
    name, _ := stringsx.TrimPrefixOK(strings.TrimSpace(rawName), "Add_")
    fmt.Println("Migration name:", name)

    // --- Run migrations (adjust to actual migrations API) ---
    // m, err := migrations.New(db, migrations.Config{Dir: "./db/migrations"})
    // if err != nil { log.Fatal(err) }
    // if err := m.Up(ctx); err != nil { log.Fatal(err) }

    fmt.Println("(migrations.Up would run here)")
    _ = ctx
}
```
