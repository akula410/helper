package mapsx_test

import (
	"sort"
	"testing"

	"github.com/akula410/helper/v2/mapsx"
)

func TestGet_Existing(t *testing.T) {
	m := map[string]int{"a": 1}
	if got := mapsx.Get(m, "a", 0); got != 1 {
		t.Fatalf("expected 1, got %d", got)
	}
}

func TestGet_Missing(t *testing.T) {
	m := map[string]int{"a": 1}
	if got := mapsx.Get(m, "b", 99); got != 99 {
		t.Fatalf("expected default 99, got %d", got)
	}
}

func TestGet_NilMap(t *testing.T) {
	var m map[string]int
	if got := mapsx.Get(m, "x", -1); got != -1 {
		t.Fatalf("expected -1, got %d", got)
	}
}

func TestHas(t *testing.T) {
	m := map[string]bool{"x": true}
	if !mapsx.Has(m, "x") {
		t.Fatal("Has should return true for existing key")
	}
	if mapsx.Has(m, "y") {
		t.Fatal("Has should return false for missing key")
	}
}

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := mapsx.Keys(m)
	sort.Strings(keys)
	if len(keys) != 3 || keys[0] != "a" || keys[1] != "b" || keys[2] != "c" {
		t.Fatalf("unexpected keys: %v", keys)
	}
}

func TestKeys_Empty(t *testing.T) {
	keys := mapsx.Keys(map[int]string{})
	if len(keys) != 0 {
		t.Fatal("expected empty slice")
	}
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1}
	vals := mapsx.Values(m)
	if len(vals) != 1 || vals[0] != 1 {
		t.Fatalf("unexpected values: %v", vals)
	}
}

func TestFindByValue_Found(t *testing.T) {
	items := []map[string]string{
		{"id": "1", "name": "alice"},
		{"id": "2", "name": "bob"},
	}
	got, ok := mapsx.FindByValue(items, "name", "bob")
	if !ok {
		t.Fatal("FindByValue should find 'bob'")
	}
	if got["id"] != "2" {
		t.Fatalf("expected id=2, got %q", got["id"])
	}
}

func TestFindByValue_NotFound(t *testing.T) {
	items := []map[string]string{{"name": "alice"}}
	_, ok := mapsx.FindByValue(items, "name", "charlie")
	if ok {
		t.Fatal("FindByValue should return false for missing value")
	}
}

func TestFindByValue_Nil(t *testing.T) {
	_, ok := mapsx.FindByValue[string](nil, "k", "v")
	if ok {
		t.Fatal("FindByValue on nil should return false")
	}
}
