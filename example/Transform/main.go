package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	var rows = []map[string]interface{}{
		{
			"group_id":"1",
			"parent_id":"0",
			"name":"Группа 1",
		},
		{
			"group_id":"2",
			"parent_id":"0",
			"name":"Группа 2",
		},
		{
			"group_id":"3",
			"parent_id":"1",
			"name":"Группа 1.1",
		},
		{
			"group_id":"4",
			"parent_id":"2",
			"name":"Группа 2.1",
		},
		{
			"group_id":"5",
			"parent_id":"3",
			"name":"Группа 1.1.1",
		},
	}
	var result = helper.Transform.LineToTree(rows, "group_id", "parent_id")

	/** Console
		 Группа 1 (1)
		- Группа 1.1 (3)
		-- Группа 1.1.1 (5)
		 Группа 2 (2)
		- Группа 2.1 (4)
	*/
	for _, r := range result {
		showTree(r, "")
	}

}

func showTree(data helper.TreeModel, p string){
	fmt.Println(p, data.Data["name"], fmt.Sprintf("(%s)", data.Id))
	if data.Children != nil {
		p += "-"
		for _, child := range data.Children {
			showTree(child, p)
		}
	}
}
