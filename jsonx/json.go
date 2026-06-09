// Package jsonx provides convenience wrappers around encoding/json.
package jsonx

import "encoding/json"

// MarshalString marshals v to a JSON string.
// Returns an error if v cannot be marshalled.
func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// MustMarshalString is like MarshalString but panics on error.
func MustMarshalString(v any) string {
	s, err := MarshalString(v)
	if err != nil {
		panic(err)
	}
	return s
}

// UnmarshalString parses the JSON string s into a value of type T.
func UnmarshalString[T any](s string) (T, error) {
	return UnmarshalBytes[T]([]byte(s))
}

// UnmarshalBytes parses the JSON bytes b into a value of type T.
func UnmarshalBytes[T any](b []byte) (T, error) {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return v, err
	}
	return v, nil
}
