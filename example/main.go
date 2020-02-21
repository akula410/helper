package main

import (
	"fmt"
	"github.com/akula410/helper"
)


func main(){
	//For work with hash BCrypt
	fmt.Println("==============helper.Hash===============")
	exampleHash("12345678")
	fmt.Println("==================end===================")



	//To reserve table identifiers
	/**
	id - current last identifier
	amount - number of records to be reserved
	 */
	fmt.Println("==============helper.ID=================")
	var id int
	id = exampleID("TABLE_NAME", "FIELD_NAME", 1)
	fmt.Println("last id = ", id)

	id = exampleID("ws_dir_manifest_ews", "ew_id", 1250)
	fmt.Println("last id = ", id)
	fmt.Println("==================end===================")


	fmt.Println("==============helper.Map=================")
	var data = map[interface{}]interface{}{"key1":"value1", "key2":"value2"}
	fmt.Println(helper.Map.Element("key1", data, "value3"))
	fmt.Println(helper.Map.Element("key2", data, "value3"))
	fmt.Println(helper.Map.Element("key3", data, "value3"))



	fmt.Println("==================end===================")


}
