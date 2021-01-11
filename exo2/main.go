package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"crypto/sha256"
	"bytes"
	"path/filepath"
)

func hashImg(path string) []byte{
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	h.Write(content)
	return h.Sum(nil)
}
func main() {

	file1_hash := hashImg("image_1.jpg")
	file2_hash := hashImg("image_2.jpg")
	file3_hash := hashImg("image_3.jpg")

	res1 := bytes.Compare(file1_hash,file2_hash)
	res2 := bytes.Compare(file1_hash,file3_hash)
	res3 := bytes.Compare(file2_hash,file3_hash)

	// if image 2 is unique
	if  res1 != 0 && res3 != 0 {
		fmt.Println("image_2 is unique")
		fmt.Println(filepath.Abs("image_2.jpg"))
	}
	// if image 3 is unique
	if res2 != 0 && res3 != 0{
		fmt.Println("image_3 is unique") 
		fmt.Println(filepath.Abs("image_3.jpg"))
	}
	// if image 1 is unique
	if res1 != 0 && res2 != 0 {
		fmt.Println("image_1 is unique") 
		fmt.Println(filepath.Abs("image_1.jpg"))
	}
}