# helper/v2

A small, safe, dependency-free utility library for Go.

```bash
go get github.com/akula410/helper/v2
```

## Sub-packages

| Package | Purpose |
|---|---|
| `password` | bcrypt hashing and verification |
| `random` | UUID v4, hex strings, opaque tokens |
| `convert` | Safe type conversions (string/int/float/bool) |
| `mapsx` | Generic map helpers |
| `slicesx` | Generic slice helpers |
| `stringsx` | String prefix/suffix utilities |
| `regexpx` | Error-returning regexp wrappers |
| `encodingx` | Base64 encode / decode |
| `jsonx` | JSON marshal / unmarshal helpers |
| `query` | URL query-string builder |
| `tree` | Flat list ↔ tree conversion |
| `idgen` | Documentation: safe ID generation patterns |

## Quick examples

```go
import (
    "github.com/akula410/helper/v2/convert"
    "github.com/akula410/helper/v2/password"
    "github.com/akula410/helper/v2/random"
)

// Password
hash, err := password.Hash("secret")
ok := password.Verify(hash, "secret")   // true

// Random
id, _    := random.UUIDv4()            // "a1b2c3d4-..."
token, _ := random.Token(32)           // 64-char hex string

// Convert
n   := convert.Int("44", 0)            // 44
f   := convert.Float64("3.14", 0)      // 3.14
b   := convert.Bool("true", false)     // true
```

## Design principles

- **No global singletons** — every function is stateless.
- **No panic in normal API** — functions return `error`; only `Must…` variants may panic.
- **Context-aware where there is I/O** — see `idgen` docs for database patterns.
- **Generics where useful** — `mapsx`, `slicesx`, `tree`, `jsonx` use type parameters.
- **Small, focused packages** — import only what you need.
- **No heavy external dependencies** — only `golang.org/x/crypto` for bcrypt.

## Integration with akula410 packages

`helper/v2` is the **bottom utility layer**. It can be used by:

```
github.com/akula410/connect
github.com/akula410/builder
github.com/akula410/migrations
```

But `helper/v2` itself **must not** import any of those packages.

Dependency direction:

```
helper/v2  ←  connect  ←  builder  ←  migrations
```

Integration examples live in [`examples/`](examples/).

## What was removed from v1

| v1 | v2 | Reason |
|---|---|---|
| Global singletons (`Hash`, `UUID`, `Map`, …) | Stateless package functions | Easier testing, no hidden state |
| `panic` on errors | Returns `error` | Predictable behaviour in production code |
| `ID.GetKey` | Removed; see `idgen/README.md` | SQL injection risk, race condition, wrong layer |
| `Transform.FileToByte` | Removed | `io.ReadAll` from stdlib is sufficient |
| `Transform.JsonToInterface` | `jsonx.UnmarshalString[T]` | Generics give type safety without reflection |

## Security notes

- Password hashes are **raw bcrypt output**, not base64-wrapped. This is the
  correct approach — bcrypt already produces a printable string.
- `Verify` never panics on malformed hashes.
- `random.*` functions use `crypto/rand` exclusively.
- `query.Encode` values are URL-escaped; keys are not interpolated into SQL.
- See `idgen/README.md` for safe database ID generation patterns.

## Testing

```bash
go test ./...
go test -race ./...
go vet ./...
```

## Versioning

This module follows [Semantic Versioning](https://semver.org).
The module path for this major version is `github.com/akula410/helper/v2`.
