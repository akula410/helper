package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/password"
)

func main() {
	hash, err := password.Hash("secret")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hash:", hash)
	fmt.Println("Verify correct:", password.Verify(hash, "secret"))
	fmt.Println("Verify wrong:  ", password.Verify(hash, "wrong"))
}
