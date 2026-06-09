package random_test

import (
	"encoding/base64"
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

// TestToken_URLSafe checks that Token returns a valid base64.RawURLEncoding string.
func TestToken_URLSafe(t *testing.T) {
	const n = 32
	s, err := random.Token(n)
	if err != nil {
		t.Fatal(err)
	}
	// base64.RawURLEncoding: no padding, URL-safe alphabet.
	// Expected length: ceil(n*4/3) = ceil(32*4/3) = 43.
	want := base64.RawURLEncoding.EncodedLen(n)
	if len(s) != want {
		t.Fatalf("expected len %d for Token(%d), got %d", want, n, len(s))
	}
	// Must decode cleanly.
	if _, err := base64.RawURLEncoding.DecodeString(s); err != nil {
		t.Fatalf("Token result is not valid base64.RawURLEncoding: %v", err)
	}
	// Must not contain padding or URL-unsafe characters.
	if strings.ContainsAny(s, "+/=") {
		t.Fatalf("Token contains URL-unsafe chars: %q", s)
	}
}

func TestToken_InvalidLength(t *testing.T) {
	_, err := random.Token(0)
	if err == nil {
		t.Fatal("expected error for length 0")
	}
	_, err = random.Token(-1)
	if err == nil {
		t.Fatal("expected error for negative length")
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
