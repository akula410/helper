# Integration example: helper/v2 + akula410/builder

This directory shows the **intended integration pattern** between `helper/v2` and
`github.com/akula410/builder`. Because the builder package API may differ between
major versions, this example is provided as documentation rather than a
compile-ready program. Adjust the builder calls to the actual API of the major
version you are using.

## Scenario

- Build URL query params with `helper/v2/query`.
- Convert incoming form/config values with `helper/v2/convert`.
- Generate a UUID for a new record with `helper/v2/random`.
- Build a parameterised SQL query with `github.com/akula410/builder`.

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    // Adjust builder import to the actual major version.
    // "github.com/akula410/builder/v2/sqlbuilder"

    "github.com/akula410/helper/v2/convert"
    "github.com/akula410/helper/v2/query"
    "github.com/akula410/helper/v2/random"
)

func main() {
    ctx := context.Background()

    // --- Build query params for an HTTP API request ---
    params, err := query.Encode(map[string]any{
        "page":   1,
        "limit":  20,
        "status": "active",
        "tags":   []string{"go", "production"},
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Query params:", params)

    // --- Convert incoming values safely ---
    page  := convert.Int("3", 1)
    limit := convert.Int("100", 20)

    // --- Generate a UUID for a new record ---
    id, err := random.UUIDv4()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("New record UUID:", id)

    // --- Build SQL with builder (adjust to actual builder API) ---
    // b := sqlbuilder.NewSelectBuilder()
    // b.Select("id", "name", "email")
    // b.From("users")
    // b.Where(b.Equal("status", "active"))
    // b.Limit(limit)
    // b.Offset((page - 1) * limit)
    // sql, args := b.Build()
    // rows, err := db.QueryContext(ctx, sql, args...)

    fmt.Printf("Would query page=%d limit=%d\n", page, limit)
    _ = ctx
}
```
