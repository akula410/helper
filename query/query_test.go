package query_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/akula410/helper/v2/query"
)

func TestEncode_Simple(t *testing.T) {
	got, err := query.Encode(map[string]any{"page": 1, "q": "hello"})
	if err != nil {
		t.Fatal(err)
	}
	// url.Values.Encode sorts keys alphabetically
	if !strings.Contains(got, "page=1") || !strings.Contains(got, "q=hello") {
		t.Fatalf("unexpected result: %q", got)
	}
}

func TestEncode_EscapeSpaces(t *testing.T) {
	got, err := query.Encode(map[string]any{"q": "hello world"})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "hello+world") && !strings.Contains(got, "hello%20world") {
		t.Fatalf("spaces not encoded in %q", got)
	}
}

func TestEncode_EscapeSpecial(t *testing.T) {
	got, err := query.Encode(map[string]any{"k": "a&b=c"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(got, "a&b=c") {
		t.Fatalf("special chars not encoded in %q", got)
	}
}

func TestEncode_StringSlice(t *testing.T) {
	got, err := query.Encode(map[string]any{"tags": []string{"go", "helper"}})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "tags=go") || !strings.Contains(got, "tags=helper") {
		t.Fatalf("slice not encoded in %q", got)
	}
}

func TestEncode_IntSlice(t *testing.T) {
	got, err := query.Encode(map[string]any{"ids": []int{1, 2, 3}})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range []string{"ids=1", "ids=2", "ids=3"} {
		if !strings.Contains(got, v) {
			t.Fatalf("missing %q in %q", v, got)
		}
	}
}

func TestEncode_Bool(t *testing.T) {
	got, err := query.Encode(map[string]any{"active": true})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "active=true") {
		t.Fatalf("unexpected: %q", got)
	}
}

func TestEncode_Float64(t *testing.T) {
	got, err := query.Encode(map[string]any{"price": 12.5})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(got, "price=12.5") {
		t.Fatalf("unexpected: %q", got)
	}
}

func TestEncode_UnsupportedType(t *testing.T) {
	_, err := query.Encode(map[string]any{"x": struct{}{}})
	if err == nil {
		t.Fatal("expected error for unsupported type")
	}
}

func TestEncodeValues(t *testing.T) {
	v := url.Values{}
	v.Set("a", "1")
	v.Set("b", "2")
	got := query.EncodeValues(v)
	if !strings.Contains(got, "a=1") || !strings.Contains(got, "b=2") {
		t.Fatalf("unexpected: %q", got)
	}
}
