package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/random"
)

func main() {
	id, err := random.UUIDv4()
	if err != nil {
		log.Fatal(err)
	}

	token, err := random.Token(32)
	if err != nil {
		log.Fatal(err)
	}

	hex16, err := random.Hex(16)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("UUID v4:", id)
	fmt.Println("Token:  ", token)
	fmt.Println("Hex16:  ", hex16)
}
