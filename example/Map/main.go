package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func main() {
	//helper.Map.Element()
	var dataElement = map[string]interface{}{"key1":"value1", "key2":"value2"}
	fmt.Println(helper.Map.Element("key1", dataElement, "value3"))
	fmt.Println(helper.Map.Element("key2", dataElement, "value3"))
	fmt.Println(helper.Map.Element("key3", dataElement, "value3"))


	//helper.Map.OfElement()
	var dataOfElement = make([]map[string]interface{}, 0)
	dataOfElement = append(dataOfElement, map[string]interface{}{"key1":"value1", "key2":"value2"})
	dataOfElement = append(dataOfElement, map[string]interface{}{"key1":"value11", "key2":"value22"})
	dataOfElement = append(dataOfElement, map[string]interface{}{"key1":"value111", "key2":"value222"})
	fmt.Println(helper.Map.OfElement("key1", "value1", dataOfElement, map[string]interface{}{"key1":"", "key2":""}))
	fmt.Println(helper.Map.OfElement("key1", "value11", dataOfElement, map[string]interface{}{"key1":"", "key2":""}))
	fmt.Println(helper.Map.OfElement("key1", "value22", dataOfElement, map[string]interface{}{"key1":"", "key2":""}))
}
