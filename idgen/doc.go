// Package idgen documents safe patterns for generating unique numeric IDs
// in Go applications backed by a relational database.
//
// # Why the v1 ID.GetKey was removed
//
// The original helper.ID.GetKey function had several design problems:
//
//  1. SQL injection risk: table and field names were interpolated directly into
//     an SQL string via fmt.Sprintf, allowing arbitrary SQL if callers passed
//     untrusted input.
//
//  2. Race condition: the in-process counter cache was not thread-safe across
//     multiple application instances (e.g., horizontal scaling).
//
//  3. Global mutable state: the Conn/ConnClose function pointers were package-
//     level variables, making testing and composition difficult.
//
//  4. Not a pure helper: ID generation is application/database logic, not a
//     utility concern that belongs in a lower-level helper library.
//
// # Recommended alternatives
//
// ## AUTO_INCREMENT / SERIAL
//
// Let the database assign the ID:
//
//	INSERT INTO users (name) VALUES (?)  -- MySQL AUTO_INCREMENT
//	id, _ := result.LastInsertId()
//
// ## UUID / ULID
//
// Use github.com/akula410/helper/v2/random:
//
//	id, err := random.UUIDv4()
//
// ## Sequence table with row-level locking
//
//	BEGIN;
//	SELECT nextval FROM sequences WHERE name = ? FOR UPDATE;
//	UPDATE sequences SET nextval = nextval + 1 WHERE name = ?;
//	COMMIT;
//
// ## Dedicated sequence service
//
// Use a centralised service (e.g., Twitter Snowflake, Sonyflake) for
// distributed, monotonically increasing IDs without database contention.
//
// # Allocator interface
//
// If you need an injectable ID allocation abstraction, implement:
//
//	type Allocator interface {
//	    Next(ctx context.Context, name string) (int64, error)
//	    Reserve(ctx context.Context, name string, count int64) (start, end int64, err error)
//	}
//
// Your implementation should:
//   - Accept context.Context for cancellation.
//   - Use database/sql with a transaction.
//   - Validate the sequence name against an allowlist before embedding it in SQL.
//   - Never use fmt.Sprintf to build SQL with user-supplied table or column names.
package idgen
