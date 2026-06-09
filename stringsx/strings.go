// Package stringsx provides additional string utility functions.
package stringsx

import "strings"

// TrimPrefixOK removes prefix from s and returns (result, true) when prefix
// is present, or (s, false) when it is not. Safe with unicode strings.
func TrimPrefixOK(s string, prefix string) (string, bool) {
	if strings.HasPrefix(s, prefix) {
		return s[len(prefix):], true
	}
	return s, false
}

// TrimSuffixOK removes suffix from s and returns (result, true) when suffix
// is present, or (s, false) when it is not. Safe with unicode strings.
func TrimSuffixOK(s string, suffix string) (string, bool) {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)], true
	}
	return s, false
}

// Empty reports whether s is the empty string.
func Empty(s string) bool { return s == "" }

// NotEmpty reports whether s is not the empty string.
func NotEmpty(s string) bool { return s != "" }
