package password_test

import (
	"testing"

	"github.com/akula410/helper/v2/password"
	"golang.org/x/crypto/bcrypt"
)

func TestHash_NotEmpty(t *testing.T) {
	h, err := password.Hash("secret")
	if err != nil {
		t.Fatal(err)
	}
	if h == "" {
		t.Fatal("hash must not be empty")
	}
}

func TestHash_EmptyPassword(t *testing.T) {
	_, err := password.Hash("")
	if err == nil {
		t.Fatal("expected error for empty password")
	}
}

func TestVerify_Valid(t *testing.T) {
	h, err := password.Hash("correct")
	if err != nil {
		t.Fatal(err)
	}
	if !password.Verify(h, "correct") {
		t.Fatal("Verify should return true for correct password")
	}
}

func TestVerify_Invalid(t *testing.T) {
	h, err := password.Hash("correct")
	if err != nil {
		t.Fatal(err)
	}
	if password.Verify(h, "wrong") {
		t.Fatal("Verify should return false for wrong password")
	}
}

func TestVerify_InvalidHash(t *testing.T) {
	if password.Verify("notahash", "secret") {
		t.Fatal("Verify should return false for invalid hash")
	}
}

func TestVerify_EmptyArgs(t *testing.T) {
	if password.Verify("", "secret") {
		t.Fatal("Verify with empty hash should return false")
	}
	if password.Verify("hash", "") {
		t.Fatal("Verify with empty plain should return false")
	}
}

func TestHashWithCost_CustomCost(t *testing.T) {
	h, err := password.HashWithCost("secret", bcrypt.MinCost)
	if err != nil {
		t.Fatal(err)
	}
	if !password.Verify(h, "secret") {
		t.Fatal("Verify should pass with custom cost hash")
	}
}
