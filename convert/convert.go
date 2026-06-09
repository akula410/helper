// Package convert provides safe type-conversion helpers.
// All functions accept any value and return a typed result, falling back to
// the caller-supplied default on failure. None of the functions panic.
package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// String converts v to a string.
// If v is already a string it is returned as-is.
// Other types are formatted with fmt.Sprintf.
// nil returns def.
func String(v any, def string) string {
	if v == nil {
		return def
	}
	switch t := v.(type) {
	case string:
		return t
	case []byte:
		return string(t)
	case bool:
		return strconv.FormatBool(t)
	case int:
		return strconv.Itoa(t)
	case int8:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(t, 10)
	case uint:
		return strconv.FormatUint(uint64(t), 10)
	case uint8:
		return strconv.FormatUint(uint64(t), 10)
	case uint16:
		return strconv.FormatUint(uint64(t), 10)
	case uint32:
		return strconv.FormatUint(uint64(t), 10)
	case uint64:
		return strconv.FormatUint(t, 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(t, 'f', -1, 64)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Int converts v to int, returning def on failure.
func Int(v any, def int) int {
	return int(Int64(v, int64(def)))
}

// Int64 converts v to int64, returning def on failure.
func Int64(v any, def int64) int64 {
	if v == nil {
		return def
	}
	switch t := v.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		if n, err := strconv.ParseInt(strings.TrimSpace(t), 10, 64); err == nil {
			return n
		}
		return def
	case []byte:
		if n, err := strconv.ParseInt(strings.TrimSpace(string(t)), 10, 64); err == nil {
			return n
		}
		return def
	default:
		return def
	}
}

// Float64 converts v to float64, returning def on failure.
func Float64(v any, def float64) float64 {
	if v == nil {
		return def
	}
	switch t := v.(type) {
	case float32:
		return float64(t)
	case float64:
		return t
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case string:
		if f, err := strconv.ParseFloat(strings.TrimSpace(t), 64); err == nil {
			return f
		}
		return def
	case []byte:
		if f, err := strconv.ParseFloat(strings.TrimSpace(string(t)), 64); err == nil {
			return f
		}
		return def
	default:
		return def
	}
}

// Bool converts v to bool, returning def on failure.
// String values "1", "t", "true", "yes", "on" (case-insensitive) are true;
// "0", "f", "false", "no", "off" are false.
func Bool(v any, def bool) bool {
	if v == nil {
		return def
	}
	switch t := v.(type) {
	case bool:
		return t
	case int, int8, int16, int32, int64:
		return Int64(t, 0) != 0
	case uint, uint8, uint16, uint32, uint64:
		return Int64(t, 0) != 0
	case string:
		switch strings.ToLower(strings.TrimSpace(t)) {
		case "1", "t", "true", "yes", "on":
			return true
		case "0", "f", "false", "no", "off":
			return false
		}
		return def
	case []byte:
		return Bool(string(t), def)
	default:
		return def
	}
}
