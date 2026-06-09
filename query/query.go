// Package query provides helpers for building URL query strings.
package query

import (
	"fmt"
	"net/url"
	"strconv"
)

// Encode converts data to a URL-encoded query string.
// Supported value types: string, int, int64, float64, bool,
// []string, []int, []int64.
// Returns an error for unsupported value types.
func Encode(data map[string]any) (string, error) {
	v := make(url.Values)
	for k, raw := range data {
		if err := addValue(v, k, raw); err != nil {
			return "", err
		}
	}
	return v.Encode(), nil
}

// EncodeValues encodes url.Values to a query string.
// This is a thin convenience wrapper around url.Values.Encode.
func EncodeValues(values url.Values) string {
	return values.Encode()
}

func addValue(v url.Values, key string, raw any) error {
	switch t := raw.(type) {
	case string:
		v.Add(key, t)
	case int:
		v.Add(key, strconv.Itoa(t))
	case int64:
		v.Add(key, strconv.FormatInt(t, 10))
	case float64:
		v.Add(key, strconv.FormatFloat(t, 'f', -1, 64))
	case bool:
		v.Add(key, strconv.FormatBool(t))
	case []string:
		for _, s := range t {
			v.Add(key, s)
		}
	case []int:
		for _, n := range t {
			v.Add(key, strconv.Itoa(n))
		}
	case []int64:
		for _, n := range t {
			v.Add(key, strconv.FormatInt(n, 10))
		}
	default:
		return fmt.Errorf("query: unsupported value type %T for key %q", raw, key)
	}
	return nil
}
