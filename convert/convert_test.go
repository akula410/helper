package convert_test

import (
	"testing"

	"github.com/akula410/helper/v2/convert"
)

// --- String ---

func TestString_FromString(t *testing.T) {
	if got := convert.String("hello", ""); got != "hello" {
		t.Fatalf("got %q", got)
	}
}

func TestString_FromInt(t *testing.T) {
	if got := convert.String(42, ""); got != "42" {
		t.Fatalf("got %q", got)
	}
}

func TestString_FromFloat(t *testing.T) {
	if got := convert.String(3.14, ""); got != "3.14" {
		t.Fatalf("got %q", got)
	}
}

func TestString_FromBool(t *testing.T) {
	if got := convert.String(true, ""); got != "true" {
		t.Fatalf("got %q", got)
	}
}

func TestString_Nil(t *testing.T) {
	if got := convert.String(nil, "default"); got != "default" {
		t.Fatalf("got %q", got)
	}
}

// --- Int ---

func TestInt_FromString(t *testing.T) {
	if got := convert.Int("44", 0); got != 44 {
		t.Fatalf("got %d", got)
	}
}

func TestInt_FromInt(t *testing.T) {
	if got := convert.Int(99, 0); got != 99 {
		t.Fatalf("got %d", got)
	}
}

func TestInt_InvalidString(t *testing.T) {
	if got := convert.Int("abc", -1); got != -1 {
		t.Fatalf("expected default -1, got %d", got)
	}
}

func TestInt_Nil(t *testing.T) {
	if got := convert.Int(nil, 7); got != 7 {
		t.Fatalf("expected default 7, got %d", got)
	}
}

func TestInt_Float(t *testing.T) {
	if got := convert.Int(3.9, 0); got != 3 {
		t.Fatalf("expected 3, got %d", got)
	}
}

// --- Int64 ---

func TestInt64_FromString(t *testing.T) {
	if got := convert.Int64("9999999999", 0); got != 9999999999 {
		t.Fatalf("got %d", got)
	}
}

// --- Float64 ---

func TestFloat64_FromString(t *testing.T) {
	if got := convert.Float64("12.50", 0); got != 12.50 {
		t.Fatalf("got %f", got)
	}
}

func TestFloat64_FromInt(t *testing.T) {
	if got := convert.Float64(5, 0); got != 5.0 {
		t.Fatalf("got %f", got)
	}
}

func TestFloat64_InvalidString(t *testing.T) {
	if got := convert.Float64("xyz", -1.0); got != -1.0 {
		t.Fatalf("expected default, got %f", got)
	}
}

func TestFloat64_Nil(t *testing.T) {
	if got := convert.Float64(nil, 0.5); got != 0.5 {
		t.Fatalf("expected 0.5, got %f", got)
	}
}

// --- Bool ---

func TestBool_FromBool(t *testing.T) {
	if !convert.Bool(true, false) {
		t.Fatal("expected true")
	}
	if convert.Bool(false, true) {
		t.Fatal("expected false")
	}
}

func TestBool_FromString_True(t *testing.T) {
	for _, s := range []any{"1", "true", "True", "TRUE", "yes", "on", "t"} {
		if !convert.Bool(s, false) {
			t.Fatalf("expected true for %q", s)
		}
	}
}

func TestBool_FromString_False(t *testing.T) {
	for _, s := range []any{"0", "false", "False", "no", "off", "f"} {
		if convert.Bool(s, true) {
			t.Fatalf("expected false for %q", s)
		}
	}
}

func TestBool_InvalidString(t *testing.T) {
	if convert.Bool("maybe", true) != true {
		t.Fatal("expected default true for unrecognised string")
	}
}

func TestBool_Nil(t *testing.T) {
	if convert.Bool(nil, true) != true {
		t.Fatal("expected default")
	}
}

func TestBool_FromInt(t *testing.T) {
	if !convert.Bool(1, false) {
		t.Fatal("expected true for 1")
	}
	if convert.Bool(0, true) {
		t.Fatal("expected false for 0")
	}
}
