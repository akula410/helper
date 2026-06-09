package jsonx_test

import (
	"testing"

	"github.com/akula410/helper/v2/jsonx"
)

func TestMarshalString(t *testing.T) {
	s, err := jsonx.MarshalString(map[string]int{"a": 1})
	if err != nil {
		t.Fatal(err)
	}
	if s != `{"a":1}` {
		t.Fatalf("got %q", s)
	}
}

func TestMarshalString_Invalid(t *testing.T) {
	// channels cannot be marshalled
	_, err := jsonx.MarshalString(make(chan int))
	if err == nil {
		t.Fatal("expected error for unmarshalable value")
	}
}

func TestMustMarshalString(t *testing.T) {
	s := jsonx.MustMarshalString([]int{1, 2, 3})
	if s != "[1,2,3]" {
		t.Fatalf("got %q", s)
	}
}

func TestMustMarshalString_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("MustMarshalString should panic on invalid value")
		}
	}()
	jsonx.MustMarshalString(make(chan int))
}

func TestUnmarshalString(t *testing.T) {
	type Point struct{ X, Y int }
	got, err := jsonx.UnmarshalString[Point](`{"X":1,"Y":2}`)
	if err != nil {
		t.Fatal(err)
	}
	if got.X != 1 || got.Y != 2 {
		t.Fatalf("unexpected: %+v", got)
	}
}

func TestUnmarshalString_Invalid(t *testing.T) {
	_, err := jsonx.UnmarshalString[map[string]any]("{not json}")
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestUnmarshalBytes(t *testing.T) {
	got, err := jsonx.UnmarshalBytes[[]int]([]byte("[1,2,3]"))
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 || got[2] != 3 {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestUnmarshalBytes_Invalid(t *testing.T) {
	_, err := jsonx.UnmarshalBytes[map[string]any]([]byte("???"))
	if err == nil {
		t.Fatal("expected error")
	}
}
