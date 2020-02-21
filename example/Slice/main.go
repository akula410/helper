package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	var data = []interface{}{"v1", "v2", "v3", "v4", "v5"}

	fmt.Println(helper.Slice.Search("v1", data))
	fmt.Println(helper.Slice.Search("v3", data))
	fmt.Println(helper.Slice.Search("v6", data))
}
