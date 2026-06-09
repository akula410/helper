package slicesx_test

import (
	"testing"

	"github.com/akula410/helper/v2/slicesx"
)

func TestContains_Int(t *testing.T) {
	if !slicesx.Contains([]int{1, 2, 3}, 2) {
		t.Fatal("expected true")
	}
	if slicesx.Contains([]int{1, 2, 3}, 9) {
		t.Fatal("expected false")
	}
}

func TestContains_String(t *testing.T) {
	if !slicesx.Contains([]string{"a", "b"}, "a") {
		t.Fatal("expected true")
	}
}

func TestContains_Nil(t *testing.T) {
	if slicesx.Contains[int](nil, 1) {
		t.Fatal("expected false for nil slice")
	}
}

func TestContains_Empty(t *testing.T) {
	if slicesx.Contains([]string{}, "x") {
		t.Fatal("expected false for empty slice")
	}
}

func TestIndex_Found(t *testing.T) {
	if i := slicesx.Index([]string{"a", "b", "c"}, "b"); i != 1 {
		t.Fatalf("expected 1, got %d", i)
	}
}

func TestIndex_NotFound(t *testing.T) {
	if i := slicesx.Index([]string{"a", "b"}, "z"); i != -1 {
		t.Fatalf("expected -1, got %d", i)
	}
}

func TestIndex_Nil(t *testing.T) {
	if i := slicesx.Index[int](nil, 0); i != -1 {
		t.Fatalf("expected -1 for nil, got %d", i)
	}
}

func TestUnique(t *testing.T) {
	got := slicesx.Unique([]int{1, 2, 2, 3, 1})
	if len(got) != 3 || got[0] != 1 || got[1] != 2 || got[2] != 3 {
		t.Fatalf("unexpected: %v", got)
	}
}

func TestUnique_NilEmpty(t *testing.T) {
	if got := slicesx.Unique[int](nil); len(got) != 0 {
		t.Fatal("expected empty")
	}
	if got := slicesx.Unique([]string{}); len(got) != 0 {
		t.Fatal("expected empty")
	}
}

func TestFilter(t *testing.T) {
	even := slicesx.Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if len(even) != 2 || even[0] != 2 || even[1] != 4 {
		t.Fatalf("unexpected: %v", even)
	}
}

func TestFilter_Nil(t *testing.T) {
	got := slicesx.Filter[int](nil, func(n int) bool { return true })
	if len(got) != 0 {
		t.Fatal("expected empty for nil input")
	}
}

func TestMap(t *testing.T) {
	doubled := slicesx.Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	if len(doubled) != 3 || doubled[0] != 2 || doubled[1] != 4 || doubled[2] != 6 {
		t.Fatalf("unexpected: %v", doubled)
	}
}

func TestMap_Empty(t *testing.T) {
	got := slicesx.Map([]string{}, func(s string) int { return 0 })
	if len(got) != 0 {
		t.Fatal("expected empty")
	}
}

type point struct{ X, Y int }

func TestContains_Struct(t *testing.T) {
	pts := []point{{1, 2}, {3, 4}}
	if !slicesx.Contains(pts, point{3, 4}) {
		t.Fatal("expected true")
	}
	if slicesx.Contains(pts, point{5, 6}) {
		t.Fatal("expected false")
	}
}
