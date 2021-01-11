package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/sha256"
)

func main() {
	// I'm reading my 'image_1.jpg' file
	content, err := ioutil.ReadFile("image_1.jpg")
	// If a have some issues
	if err != nil {
		log.Fatal(err)
	}
	// else: I convert my byte in SHA256
	h := sha256.New()
	h.Write(content)
	fmt.Printf("File contents: %x", h.Sum(nil))
}