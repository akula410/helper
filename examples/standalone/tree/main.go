package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/tree"
)

type Category struct {
	ID       int
	ParentID *int
	Name     string
}

func ptr(n int) *int { return &n }

func main() {
	items := []Category{
		{1, nil, "Electronics"},
		{2, ptr(1), "Phones"},
		{3, ptr(1), "Laptops"},
		{4, ptr(2), "Android"},
		{5, ptr(2), "iOS"},
	}

	roots, err := tree.FromFlat(
		items,
		func(c Category) int { return c.ID },
		func(c Category) *int { return c.ParentID },
	)
	if err != nil {
		log.Fatal(err)
	}

	printTree(roots, 0)

	fmt.Println("\nFlattened:")
	for _, c := range tree.Flatten(roots) {
		fmt.Println(" -", c.Name)
	}
}

func printTree(nodes []*tree.Node[Category, int], depth int) {
	indent := ""
	for range depth {
		indent += "  "
	}
	for _, n := range nodes {
		fmt.Printf("%s%s (id=%d)\n", indent, n.Data.Name, n.ID)
		printTree(n.Children, depth+1)
	}
}
