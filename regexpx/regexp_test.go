package regexpx_test

import (
	"testing"

	"github.com/akula410/helper/v2/regexpx"
)

func TestFind_Valid(t *testing.T) {
	got, err := regexpx.Find(`\d+`, "abc 123 def")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 || got[0] != "123" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestFind_NoMatch(t *testing.T) {
	got, err := regexpx.Find(`\d+`, "abc")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestFind_InvalidPattern(t *testing.T) {
	_, err := regexpx.Find(`[invalid`, "text")
	if err == nil {
		t.Fatal("expected error for invalid regexp")
	}
}

func TestFindString_Found(t *testing.T) {
	s, ok, err := regexpx.FindString(`go\d+`, "go123")
	if err != nil || !ok || s != "go123" {
		t.Fatalf("got %q, ok=%v, err=%v", s, ok, err)
	}
}

func TestFindString_NotFound(t *testing.T) {
	s, ok, err := regexpx.FindString(`go\d+`, "python")
	if err != nil || ok || s != "" {
		t.Fatalf("got %q, ok=%v, err=%v", s, ok, err)
	}
}

func TestFindAll_Valid(t *testing.T) {
	got, err := regexpx.FindAll(`\d+`, "a1 b22 c333")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Fatalf("expected 3 matches, got %d: %v", len(got), got)
	}
}

func TestFindAll_NoMatch(t *testing.T) {
	got, err := regexpx.FindAll(`\d+`, "abc")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 0 {
		t.Fatalf("expected empty, got %v", got)
	}
}

func TestFindAll_InvalidPattern(t *testing.T) {
	_, err := regexpx.FindAll(`[bad`, "text")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestMatch_Valid(t *testing.T) {
	ok, err := regexpx.Match(`^\d+$`, "12345")
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestMatch_Invalid(t *testing.T) {
	_, err := regexpx.Match(`[bad`, "text")
	if err == nil {
		t.Fatal("expected error for invalid pattern")
	}
}

func TestMustFind_Valid(t *testing.T) {
	got := regexpx.MustFind(`\w+`, "hello")
	if len(got) == 0 || got[0] != "hello" {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestMustFind_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustFind should panic on invalid pattern")
		}
	}()
	regexpx.MustFind(`[bad`, "text")
}
