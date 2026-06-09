// Package mapsx provides generic map helper functions.
package mapsx

// Get returns the value for key in m, or def if the key is absent.
func Get[K comparable, V any](m map[K]V, key K, def V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return def
}

// Has reports whether m contains key.
func Has[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// Keys returns all keys of m in an unspecified order.
func Keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// Values returns all values of m in an unspecified order.
func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

// FindByValue searches items for the first element where items[i][key] == value.
// Returns (element, true) when found, or (zero, false) when not found.
func FindByValue[V comparable](items []map[string]V, key string, value V) (map[string]V, bool) {
	for _, item := range items {
		if v, ok := item[key]; ok && v == value {
			return item, true
		}
	}
	return nil, false
}
