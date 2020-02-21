package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	var string1 = "Hello world !!!"
	fmt.Println(helper.Reg.Find(`\s+[a-z]+\s+`, string1))

	fmt.Println(helper.Reg.FindAll(`\s*[a-zA-Z]+\s+`, string1))
}
