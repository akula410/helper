package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/query"
)

func main() {
	result, err := query.Encode(map[string]any{
		"page":  1,
		"limit": 50,
		"q":     "hello world",
		"tags":  []string{"go", "helper"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
