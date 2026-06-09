// This example demonstrates the intended integration style between
// helper/v2 and github.com/akula410/migrations.
//
// Adjust migrations API calls to the actual github.com/akula410/migrations major version.
package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/akula410/helper/v2/convert"
	"github.com/akula410/helper/v2/random"
	"github.com/akula410/helper/v2/stringsx"
)

func main() {
	// --- Build a migration label from env/config values ---
	rawVersion := "  1  "
	rawEnv := "STAGING"

	version := convert.Int(rawVersion, 0)
	env := strings.ToLower(strings.TrimSpace(rawEnv))
	fmt.Printf("Running migrations for env=%s version=%d\n", env, version)

	// --- Generate a migration-run ID for logging/auditing ---
	runID, err := random.UUIDv4()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Migration run ID:", runID)

	// --- Normalise a migration name ---
	rawName := "  Add_Users_Table  "
	name, _ := stringsx.TrimPrefixOK(strings.TrimSpace(rawName), "Add_")
	fmt.Println("Migration name:", name)

	// --- This example demonstrates the intended integration style. ---
	// Adjust migrations API calls to the actual github.com/akula410/migrations major version.
	//
	// Example (adapt to actual migrations API):
	//
	//   m, err := migrations.New(db, migrations.Config{
	//       Dir: "./db/migrations",
	//   })
	//   if err != nil { log.Fatal(err) }
	//
	//   if err := m.Up(ctx); err != nil { log.Fatal(err) }

	fmt.Println("(migrations.Up would run here)")
}
