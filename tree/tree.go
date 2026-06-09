// Package tree converts flat lists into tree structures and back.
//
// It supports multiple roots, handles missing parents gracefully,
// and detects cycles using depth-first traversal.
package tree

import "fmt"

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
// Returns ErrCycle if a cycle is detected.
func FromFlat[T any, ID comparable](
	items []T,
	getID func(T) ID,
	getParentID func(T) *ID,
) ([]*Node[T, ID], error) {
	nodes := make(map[ID]*Node[T, ID], len(items))
	for _, item := range items {
		id := getID(item)
		nodes[id] = &Node[T, ID]{
			ID:       id,
			ParentID: getParentID(item),
			Data:     item,
			Children: []*Node[T, ID]{},
		}
	}

	var roots []*Node[T, ID]
	for _, n := range nodes {
		if n.ParentID == nil {
			roots = append(roots, n)
		} else if parent, ok := nodes[*n.ParentID]; ok {
			parent.Children = append(parent.Children, n)
		} else {
			// parent not found — treat as root
			roots = append(roots, n)
		}
	}

	if err := detectCycles(roots); err != nil {
		return nil, err
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

// detectCycles walks all nodes reachable from roots and returns an error if
// any node is visited more than once (which would indicate a cycle in the data).
func detectCycles[T any, ID comparable](roots []*Node[T, ID]) error {
	visited := make(map[ID]bool)
	for _, r := range roots {
		if err := dfsCheck(r, visited); err != nil {
			return err
		}
	}
	return nil
}

func dfsCheck[T any, ID comparable](n *Node[T, ID], visited map[ID]bool) error {
	if visited[n.ID] {
		return fmt.Errorf("tree: cycle detected at node %v", n.ID)
	}
	visited[n.ID] = true
	for _, child := range n.Children {
		if err := dfsCheck(child, visited); err != nil {
			return err
		}
	}
	return nil
}
