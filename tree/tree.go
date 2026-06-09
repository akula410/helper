// Package tree converts flat lists into tree structures and back.
//
// It supports multiple roots, handles missing parents gracefully,
// detects cycles even when no root nodes exist, and rejects duplicate IDs.
package tree

import (
	"errors"
	"fmt"
)

// ErrCycle is returned by FromFlat when a cycle is detected in the parent chain.
var ErrCycle = errors.New("tree: cycle detected")

// ErrDuplicateID is returned by FromFlat when two items share the same ID.
var ErrDuplicateID = errors.New("tree: duplicate ID")

// Node wraps a data item with its tree relationships.
type Node[T any, ID comparable] struct {
	ID       ID
	ParentID *ID
	Data     T
	Children []*Node[T, ID]
}

// FromFlat builds a forest (slice of root nodes) from a flat slice of items.
//
// getID extracts the unique identifier from an item.
// getParentID returns a pointer to the parent ID, or nil for root items.
//
// Items whose parent is not found in the list are treated as roots.
// Returns ErrDuplicateID if two items share the same ID.
// Returns ErrCycle if a cycle is detected in the parent chain.
// Cycle detection works correctly even when no root nodes exist.
func FromFlat[T any, ID comparable](
	items []T,
	getID func(T) ID,
	getParentID func(T) *ID,
) ([]*Node[T, ID], error) {
	nodes := make(map[ID]*Node[T, ID], len(items))

	// Phase 1: build node map, reject duplicates.
	for _, item := range items {
		id := getID(item)
		if _, exists := nodes[id]; exists {
			return nil, fmt.Errorf("%w: id=%v", ErrDuplicateID, id)
		}
		nodes[id] = &Node[T, ID]{
			ID:       id,
			ParentID: getParentID(item),
			Data:     item,
			Children: []*Node[T, ID]{},
		}
	}

	// Phase 2: detect cycles by walking each node's parent chain.
	// Uses white/gray/black colouring — works even when no root nodes exist.
	if err := detectCycles(nodes); err != nil {
		return nil, err
	}

	// Phase 3: link children to parents, collect roots.
	var roots []*Node[T, ID]
	for _, n := range nodes {
		if n.ParentID == nil {
			roots = append(roots, n)
		} else if parent, ok := nodes[*n.ParentID]; ok {
			parent.Children = append(parent.Children, n)
		} else {
			// Parent not found in the list — treat node as a root.
			roots = append(roots, n)
		}
	}

	return roots, nil
}

// Flatten performs a depth-first traversal and returns the data values of all
// nodes in the forest as a flat slice.
func Flatten[T any, ID comparable](roots []*Node[T, ID]) []T {
	out := make([]T, 0)
	for _, r := range roots {
		flattenNode(r, &out)
	}
	return out
}

func flattenNode[T any, ID comparable](n *Node[T, ID], out *[]T) {
	*out = append(*out, n.Data)
	for _, child := range n.Children {
		flattenNode(child, out)
	}
}

// detectCycles walks each node's parent chain using DFS with white/gray/black
// colouring. A node coloured gray means it is on the current path — reaching
// it again indicates a cycle.
func detectCycles[T any, ID comparable](nodes map[ID]*Node[T, ID]) error {
	const (
		white = 0 // not visited
		gray  = 1 // currently on the path
		black = 2 // fully resolved
	)
	color := make(map[ID]int, len(nodes))

	var visit func(id ID) error
	visit = func(id ID) error {
		switch color[id] {
		case black:
			return nil
		case gray:
			return fmt.Errorf("%w: node %v", ErrCycle, id)
		}
		color[id] = gray
		n := nodes[id]
		if n.ParentID != nil {
			if _, exists := nodes[*n.ParentID]; exists {
				if err := visit(*n.ParentID); err != nil {
					return err
				}
			}
		}
		color[id] = black
		return nil
	}

	for id := range nodes {
		if err := visit(id); err != nil {
			return err
		}
	}
	return nil
}
