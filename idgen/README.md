# idgen

The `idgen` package contains **documentation only** — no executable code.

It explains why the v1 `ID.GetKey` function was removed and recommends safe
alternatives for generating unique numeric IDs in Go applications.

## Why `ID.GetKey` was removed

| Problem | Detail |
|---|---|
| SQL injection | Table and field names were built with `fmt.Sprintf`, allowing arbitrary SQL if callers passed untrusted input |
| Race condition | The in-process counter cache fails under concurrent writes from multiple instances |
| Global state | `Conn` / `ConnClose` were package-level mutable variables, making testing difficult |
| Wrong layer | ID generation is application/database logic, not a utility concern |

## Safe alternatives

### 1. AUTO_INCREMENT / SERIAL (recommended)

Let the database assign the ID:

```go
result, err := db.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", name)
id, _ := result.LastInsertId()
```

### 2. UUID v4

```go
import "github.com/akula410/helper/v2/random"

id, err := random.UUIDv4()
```

### 3. Sequence table with row-level locking

```sql
BEGIN;
SELECT nextval FROM sequences WHERE name = ? FOR UPDATE;
UPDATE sequences SET nextval = nextval + 1 WHERE name = ?;
COMMIT;
```

```go
func nextID(ctx context.Context, db *sql.DB, seqName string) (int64, error) {
    // seqName MUST be validated against a known allowlist before use.
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return 0, err
    }
    defer tx.Rollback()

    var next int64
    err = tx.QueryRowContext(ctx,
        "SELECT nextval FROM sequences WHERE name = ? FOR UPDATE", seqName,
    ).Scan(&next)
    if err != nil {
        return 0, err
    }
    _, err = tx.ExecContext(ctx,
        "UPDATE sequences SET nextval = nextval + 1 WHERE name = ?", seqName,
    )
    if err != nil {
        return 0, err
    }
    return next, tx.Commit()
}
```

### 4. Allocator interface

If you need a pluggable abstraction:

```go
type Allocator interface {
    Next(ctx context.Context, name string) (int64, error)
    Reserve(ctx context.Context, name string, count int64) (start, end int64, err error)
}
```

Your implementation **must**:
- Accept `context.Context`.
- Use a transaction.
- Validate `name` against an allowlist — never embed it in SQL directly.
- Return `error` instead of panicking.
