package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/jsonx"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	u := User{Name: "Alice", Email: "alice@example.com", Age: 30}

	s, err := jsonx.MarshalString(u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON:", s)

	u2, err := jsonx.UnmarshalString[User](s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Parsed: %+v\n", u2)

	nums, err := jsonx.UnmarshalBytes[[]int]([]byte("[1,2,3,4,5]"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Nums:", nums)
}
