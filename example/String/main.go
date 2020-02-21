package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	var rout = "word1/word2/word3"

	fmt.Println(helper.String.DeleteStart(rout, "word1"))

	fmt.Println(helper.String.DeleteEnd(rout, "word1"))

	fmt.Println(helper.String.DeleteEnd(rout, "word3"))
}
