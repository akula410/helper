package tree_test

import (
	"errors"
	"testing"

	"github.com/akula410/helper/v2/tree"
)

type cat struct {
	ID       int
	ParentID *int
	Name     string
}

func ptr(n int) *int { return &n }

func getID(c cat) int      { return c.ID }
func getParent(c cat) *int { return c.ParentID }

func TestFromFlat_Simple(t *testing.T) {
	items := []cat{
		{1, nil, "Root"},
		{2, ptr(1), "Child1"},
		{3, ptr(1), "Child2"},
		{4, ptr(2), "GrandChild"},
	}
	roots, err := tree.FromFlat(items, getID, getParent)
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
	roots, err := tree.FromFlat(items, getID, getParent)
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 2 {
		t.Fatalf("expected 2 roots, got %d", len(roots))
	}
}

func TestFromFlat_MissingParent(t *testing.T) {
	// Item 2 points to parent 99 which does not exist — treated as root.
	items := []cat{
		{1, nil, "Root"},
		{2, ptr(99), "Orphan"},
	}
	roots, err := tree.FromFlat(items, getID, getParent)
	if err != nil {
		t.Fatal(err)
	}
	if len(roots) != 2 {
		t.Fatalf("expected 2 roots (orphan treated as root), got %d", len(roots))
	}
}

func TestFromFlat_Empty(t *testing.T) {
	roots, err := tree.FromFlat([]cat{}, getID, getParent)
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
	roots, err := tree.FromFlat(items, getID, getParent)
	if err != nil {
		t.Fatal(err)
	}
	flat := tree.Flatten(roots)
	if len(flat) != 2 {
		t.Fatalf("expected 2 items, got %d", len(flat))
	}
}

// TestFromFlat_Cycle_Direct tests detection of a direct 1→2→1 cycle.
// Item 1 has parent 2, item 2 has parent 1 — neither has a nil parent,
// so there are no root nodes. The cycle must still be detected.
func TestFromFlat_Cycle_Direct(t *testing.T) {
	items := []cat{
		{1, ptr(2), "A"},
		{2, ptr(1), "B"},
	}
	_, err := tree.FromFlat(items, getID, getParent)
	if err == nil {
		t.Fatal("expected error for 1→2→1 cycle, got nil")
	}
	if !errors.Is(err, tree.ErrCycle) {
		t.Fatalf("expected ErrCycle, got: %v", err)
	}
}

// TestFromFlat_Cycle_NoRoots tests that cycle detection works even when
// all nodes participate in a cycle (no root nodes exist).
func TestFromFlat_Cycle_NoRoots(t *testing.T) {
	// 1→2→3→1: three-node cycle with no roots.
	items := []cat{
		{1, ptr(2), "A"},
		{2, ptr(3), "B"},
		{3, ptr(1), "C"},
	}
	_, err := tree.FromFlat(items, getID, getParent)
	if err == nil {
		t.Fatal("expected ErrCycle for three-node cycle without roots")
	}
	if !errors.Is(err, tree.ErrCycle) {
		t.Fatalf("expected ErrCycle, got: %v", err)
	}
}

// TestFromFlat_DuplicateID tests that duplicate IDs return ErrDuplicateID.
func TestFromFlat_DuplicateID(t *testing.T) {
	items := []cat{
		{1, nil, "A"},
		{1, nil, "A_dup"},
	}
	_, err := tree.FromFlat(items, getID, getParent)
	if err == nil {
		t.Fatal("expected error for duplicate IDs, got nil")
	}
	if !errors.Is(err, tree.ErrDuplicateID) {
		t.Fatalf("expected ErrDuplicateID, got: %v", err)
	}
}
