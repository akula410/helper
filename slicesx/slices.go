// Package slicesx provides generic slice helper functions.
package slicesx

// Contains reports whether value is present in items.
func Contains[T comparable](items []T, value T) bool {
	for _, v := range items {
		if v == value {
			return true
		}
	}
	return false
}

// Index returns the first index of value in items, or -1 if not found.
func Index[T comparable](items []T, value T) int {
	for i, v := range items {
		if v == value {
			return i
		}
	}
	return -1
}

// Unique returns a new slice with duplicate values removed, preserving order.
func Unique[T comparable](items []T) []T {
	seen := make(map[T]struct{}, len(items))
	out := make([]T, 0, len(items))
	for _, v := range items {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			out = append(out, v)
		}
	}
	return out
}

// Filter returns a new slice containing only elements for which fn returns true.
func Filter[T any](items []T, fn func(T) bool) []T {
	out := make([]T, 0, len(items))
	for _, v := range items {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

// Map applies fn to each element and returns the resulting slice.
func Map[T any, R any](items []T, fn func(T) R) []R {
	out := make([]R, len(items))
	for i, v := range items {
		out[i] = fn(v)
	}
	return out
}
