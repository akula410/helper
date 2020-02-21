package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	/** helper.UUID.GetUUID() */
	fmt.Println(helper.UUID.GetUUID())
	fmt.Println(helper.UUID.GetUUID("*"))
	fmt.Println(helper.UUID.GetUUID())

	/** helper.UUID.GetUniqName() */
	fmt.Println(helper.UUID.GetUniqName(2))
	fmt.Println(helper.UUID.GetUniqName(4))
	fmt.Println(helper.UUID.GetUniqName(8))
}
