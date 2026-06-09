package random_test

import (
	"regexp"
	"strings"
	"testing"

	"github.com/akula410/helper/v2/random"
)

var uuidRe = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

func TestUUIDv4_Format(t *testing.T) {
	id, err := random.UUIDv4()
	if err != nil {
		t.Fatal(err)
	}
	if !uuidRe.MatchString(id) {
		t.Fatalf("UUIDv4 format invalid: %q", id)
	}
}

func TestUUIDv4_Version(t *testing.T) {
	for range 50 {
		id, err := random.UUIDv4()
		if err != nil {
			t.Fatal(err)
		}
		parts := strings.Split(id, "-")
		if len(parts) != 5 {
			t.Fatalf("expected 5 parts, got %d", len(parts))
		}
		if parts[2][0] != '4' {
			t.Fatalf("version bit wrong in %q", id)
		}
		v := parts[3][0]
		if v != '8' && v != '9' && v != 'a' && v != 'b' {
			t.Fatalf("variant bit wrong in %q", id)
		}
	}
}

func TestUUIDv4_Uniqueness(t *testing.T) {
	seen := make(map[string]struct{}, 1000)
	for range 1000 {
		id, err := random.UUIDv4()
		if err != nil {
			t.Fatal(err)
		}
		if _, dup := seen[id]; dup {
			t.Fatalf("duplicate UUID: %q", id)
		}
		seen[id] = struct{}{}
	}
}

func TestMustUUIDv4(t *testing.T) {
	id := random.MustUUIDv4()
	if !uuidRe.MatchString(id) {
		t.Fatalf("MustUUIDv4 format invalid: %q", id)
	}
}

func TestHex_Length(t *testing.T) {
	s, err := random.Hex(16)
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 32 {
		t.Fatalf("expected len 32, got %d", len(s))
	}
}

func TestHex_InvalidLength(t *testing.T) {
	_, err := random.Hex(0)
	if err == nil {
		t.Fatal("expected error for length 0")
	}
	_, err = random.Hex(-1)
	if err == nil {
		t.Fatal("expected error for negative length")
	}
}

func TestMustHex_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustHex should panic on invalid length")
		}
	}()
	random.MustHex(0)
}

func TestToken_Length(t *testing.T) {
	s, err := random.Token(32)
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 64 {
		t.Fatalf("expected len 64, got %d", len(s))
	}
}

func TestMustToken_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustToken should panic on invalid length")
		}
	}()
	random.MustToken(-5)
}
