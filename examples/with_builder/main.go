// This example demonstrates the intended integration style between
// helper/v2 and github.com/akula410/builder.
//
// Adjust builder API calls to the actual github.com/akula410/builder major version.
package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/convert"
	"github.com/akula410/helper/v2/query"
	"github.com/akula410/helper/v2/random"
)

func main() {
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

	// --- Convert incoming form/config values safely ---
	rawPage := "3"
	rawLimit := "100"
	page := convert.Int(rawPage, 1)
	limit := convert.Int(rawLimit, 20)
	fmt.Printf("Fetching page=%d limit=%d\n", page, limit)

	// --- Generate a UUID for a new record ---
	id, err := random.UUIDv4()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New record ID:", id)

	// --- Build SQL with builder (placeholder) ---
	// This example demonstrates the intended integration style.
	// Adjust builder API calls to the actual github.com/akula410/builder major version.
	//
	// Example (adapt to actual builder API):
	//   b := sqlbuilder.NewSelectBuilder()
	//   b.Select("id", "name", "email")
	//   b.From("users")
	//   b.Where(b.Equal("status", "active"))
	//   b.Limit(limit)
	//   b.Offset((page - 1) * limit)
	//   sql, args := b.Build()
	//   rows, err := db.QueryContext(ctx, sql, args...)

	fmt.Println("(builder query would be constructed here)")
	fmt.Println("page:", page, "limit:", limit, "id:", id)
}
