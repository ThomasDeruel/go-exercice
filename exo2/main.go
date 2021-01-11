package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/sha256"
)

func main() {
	content, err := ioutil.ReadFile("image_1.jpg")
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write(content)
	fmt.Printf("File contents: %x", h.Sum(nil))
}