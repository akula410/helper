package encodingx_test

import (
	"testing"

	"github.com/akula410/helper/v2/encodingx"
)

func TestBase64Encode(t *testing.T) {
	got := encodingx.Base64Encode([]byte("hello"))
	if got != "aGVsbG8=" {
		t.Fatalf("expected aGVsbG8=, got %q", got)
	}
}

func TestBase64Decode_Valid(t *testing.T) {
	b, err := encodingx.Base64Decode("aGVsbG8=")
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != "hello" {
		t.Fatalf("expected 'hello', got %q", string(b))
	}
}

func TestBase64Decode_Invalid(t *testing.T) {
	_, err := encodingx.Base64Decode("not_valid_base64!!!")
	if err == nil {
		t.Fatal("expected error for invalid base64")
	}
}

func TestBase64RoundTrip(t *testing.T) {
	data := []byte("binary\x00data\xff")
	encoded := encodingx.Base64Encode(data)
	decoded, err := encodingx.Base64Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}
	if string(decoded) != string(data) {
		t.Fatal("round-trip mismatch")
	}
}

func TestMustBase64Decode_Valid(t *testing.T) {
	b := encodingx.MustBase64Decode("aGVsbG8=")
	if string(b) != "hello" {
		t.Fatalf("expected 'hello', got %q", string(b))
	}
}

func TestMustBase64Decode_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustBase64Decode should panic on invalid input")
		}
	}()
	encodingx.MustBase64Decode("!!invalid!!")
}
