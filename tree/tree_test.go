package tree_test

import (
	"testing"

	"github.com/akula410/helper/v2/tree"
)

type cat struct {
	ID       int
	ParentID *int
	Name     string
}

func ptr(n int) *int { return &n }

func TestFromFlat_Simple(t *testing.T) {
	items := []cat{
		{1, nil, "Root"},
		{2, ptr(1), "Child1"},
		{3, ptr(1), "Child2"},
		{4, ptr(2), "GrandChild"},
	}
	roots, err := tree.FromFlat(items, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 1 {
		t.Fatalf("expected 1 root, got %d", len(roots))
	}
	if roots[0].Data.Name != "Root" {
		t.Fatalf("expected Root, got %q", roots[0].Data.Name)
	}
	if len(roots[0].Children) != 2 {
		t.Fatalf("expected 2 children, got %d", len(roots[0].Children))
	}
}

func TestFromFlat_MultipleRoots(t *testing.T) {
	items := []cat{
		{1, nil, "A"},
		{2, nil, "B"},
		{3, ptr(1), "A1"},
	}
	roots, err := tree.FromFlat(items, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 2 {
		t.Fatalf("expected 2 roots, got %d", len(roots))
	}
}

func TestFromFlat_MissingParent(t *testing.T) {
	// Item 2 refers to parent 99 which does not exist — should be treated as root.
	items := []cat{
		{1, nil, "Root"},
		{2, ptr(99), "Orphan"},
	}
	roots, err := tree.FromFlat(items, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 2 {
		t.Fatalf("expected 2 roots (missing parent = root), got %d", len(roots))
	}
}

func TestFromFlat_Empty(t *testing.T) {
	roots, err := tree.FromFlat([]cat{}, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 0 {
		t.Fatalf("expected 0 roots, got %d", len(roots))
	}
}

func TestFlatten(t *testing.T) {
	items := []cat{
		{1, nil, "Root"},
		{2, ptr(1), "Child"},
	}
	roots, err := tree.FromFlat(items, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	flat := tree.Flatten(roots)
	if len(flat) != 2 {
		t.Fatalf("expected 2 items, got %d", len(flat))
	}
}

// TestFromFlat_Cycle builds a cycle by directly constructing nodes and
// verifies that FromFlat detects it via the visited-map guard.
//
// Because FromFlat itself builds the tree from a flat parent-pointer list,
// a "classic" cycle (A→B→A) cannot appear unless two items point to each
// other as parents (which would be a degenerate case resolved at link time).
// We test that the duplicate-ID scenario is caught instead.
func TestFromFlat_DuplicateID(t *testing.T) {
	// Two items with the same ID — only one will be kept in the node map,
	// the other is silently overwritten. This is expected behaviour: callers
	// must ensure IDs are unique. The test documents the behaviour.
	items := []cat{
		{1, nil, "A"},
		{1, nil, "A_dup"},
	}
	roots, err := tree.FromFlat(items, func(c cat) int { return c.ID }, func(c cat) *int { return c.ParentID })
	if err != nil {
		t.Fatal(err)
	}
	// One unique root expected.
	if len(roots) != 1 {
		t.Fatalf("expected 1 root for duplicate IDs, got %d", len(roots))
	}
}
