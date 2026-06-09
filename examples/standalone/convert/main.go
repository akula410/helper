package main

import (
	"fmt"

	"github.com/akula410/helper/v2/convert"
)

func main() {
	age := convert.Int("44", 0)
	active := convert.Bool("true", false)
	price := convert.Float64("12.50", 0)
	label := convert.String(age, "")

	fmt.Println("age:   ", age)
	fmt.Println("active:", active)
	fmt.Println("price: ", price)
	fmt.Println("label: ", label)

	fmt.Println("nil→int:   ", convert.Int(nil, -1))
	fmt.Println("bad→float: ", convert.Float64("not_a_number", 0.0))
}
