package stringsx_test

import (
	"testing"

	"github.com/akula410/helper/v2/stringsx"
)

func TestTrimPrefixOK_Found(t *testing.T) {
	got, ok := stringsx.TrimPrefixOK("hello world", "hello ")
	if !ok || got != "world" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestTrimPrefixOK_NotFound(t *testing.T) {
	got, ok := stringsx.TrimPrefixOK("hello", "bye")
	if ok || got != "hello" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestTrimPrefixOK_Empty(t *testing.T) {
	got, ok := stringsx.TrimPrefixOK("", "x")
	if ok || got != "" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestTrimSuffixOK_Found(t *testing.T) {
	got, ok := stringsx.TrimSuffixOK("hello.go", ".go")
	if !ok || got != "hello" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestTrimSuffixOK_NotFound(t *testing.T) {
	got, ok := stringsx.TrimSuffixOK("hello.go", ".ts")
	if ok || got != "hello.go" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestEmpty(t *testing.T) {
	if !stringsx.Empty("") {
		t.Fatal("expected true for empty string")
	}
	if stringsx.Empty("x") {
		t.Fatal("expected false for non-empty string")
	}
}

func TestNotEmpty(t *testing.T) {
	if !stringsx.NotEmpty("hi") {
		t.Fatal("expected true")
	}
	if stringsx.NotEmpty("") {
		t.Fatal("expected false for empty string")
	}
}

func TestTrimPrefixOK_Unicode(t *testing.T) {
	got, ok := stringsx.TrimPrefixOK("Привет мир", "Привет ")
	if !ok || got != "мир" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}

func TestTrimSuffixOK_Unicode(t *testing.T) {
	got, ok := stringsx.TrimSuffixOK("файл.txt", ".txt")
	if !ok || got != "файл" {
		t.Fatalf("got %q, ok=%v", got, ok)
	}
}
