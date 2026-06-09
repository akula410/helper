// Package random provides cryptographically secure random generators:
// UUID v4, hex strings, and opaque tokens.
package random

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// ErrInvalidLength is returned when a non-positive length is requested.
var ErrInvalidLength = errors.New("random: length must be greater than zero")

// UUIDv4 generates a random UUID version 4 string in standard 8-4-4-4-12 format.
// Uses crypto/rand as the entropy source.
func UUIDv4() (string, error) {
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	// Set version bits to 4 (0100xxxx) in octet 6.
	b[6] = (b[6] & 0x0f) | 0x40
	// Set variant bits to 10xxxxxx in octet 8 (RFC 4122).
	b[8] = (b[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16]), nil
}

// MustUUIDv4 is like UUIDv4 but panics on error.
func MustUUIDv4() string {
	v, err := UUIDv4()
	if err != nil {
		panic(err)
	}
	return v
}

// Hex returns a random hex-encoded string of exactly n bytes (resulting string length = 2*n).
func Hex(n int) (string, error) {
	if n <= 0 {
		return "", ErrInvalidLength
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// MustHex is like Hex but panics on error.
func MustHex(n int) string {
	s, err := Hex(n)
	if err != nil {
		panic(err)
	}
	return s
}

// Token returns a URL-safe base64 (hex) random string of n random bytes.
// The returned string has length 2*n.
func Token(n int) (string, error) {
	return Hex(n)
}

// MustToken is like Token but panics on error.
func MustToken(n int) string {
	s, err := Token(n)
	if err != nil {
		panic(err)
	}
	return s
}
