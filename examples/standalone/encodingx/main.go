package main

import (
	"fmt"
	"log"

	"github.com/akula410/helper/v2/encodingx"
)

func main() {
	original := []byte("hello, base64!")
	encoded := encodingx.Base64Encode(original)
	fmt.Println("Encoded:", encoded)

	decoded, err := encodingx.Base64Decode(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Decoded:", string(decoded))

	_, err = encodingx.Base64Decode("!!!invalid!!!")
	if err != nil {
		fmt.Println("Error (expected):", err)
	}
}
