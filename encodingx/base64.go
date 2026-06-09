// Package encodingx provides helpers for encoding operations.
// Currently supports standard base64 encoding and decoding.
package encodingx

import (
	"encoding/base64"
)

// Base64Encode encodes data to a standard base64 string.
func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode decodes a standard base64 string.
// Returns an error for invalid input; never panics.
func Base64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// MustBase64Decode is like Base64Decode but panics on error.
func MustBase64Decode(s string) []byte {
	b, err := Base64Decode(s)
	if err != nil {
		panic(err)
	}
	return b
}
