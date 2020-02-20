package main

import (
	"fmt"
	"github.com/akula410/helper"
)

func exampleHash(pass string){
	fmt.Println("password: ", pass)
	var hash1 = helper.Hash.GetHashBCrypt(pass)
	var hash2 = helper.Hash.GetHashBCrypt(pass)

	fmt.Println("Hash 1 for password: ", hash1)
	fmt.Println("Hash 2 for password: ", hash2)

	fmt.Println("Result hash 1 for password: ", helper.Hash.AskHashBCrypt(hash1, pass))
	fmt.Println("Result hash 2 for password: ", helper.Hash.AskHashBCrypt(hash2, pass))
}
