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
	// we can defined var like this too
	// var myVar mytype = someValue_withtype
	
	all_info := make(map[string][]byte)

	var (
		files []string= []string{"image_1.jpg","image_2.jpg","image_3.jpg"}
		files_hash [][]byte = [][]byte{}
	)
	for _,file := range files {
		f_hash := hashImg(file)
		files_hash = append(files_hash,f_hash)
		all_info[file] = f_hash
	}

	// find indes (duplicates byte)
	indices := []int{}
	for index,file_hash := range files_hash {
		others_files_hash := files_hash[index+1:]
		if len(others_files_hash) < 0 {
			continue
		}
		for index2,file_hash_two := range others_files_hash{
			res := bytes.Compare(file_hash,file_hash_two)
			if res == 0 {
				indices = append(indices,index,index2)
			}	
		}
	}
	// remove duplicate values
	for i := range indices {
		copy(files_hash[i:], files_hash[i+1:])
		files_hash[len(files_hash)-1] = []byte{}  
		files_hash = files_hash[:len(files_hash)-1] 
	}
	
	for text,info := range all_info {
		for _,file := range files_hash {
			if bytes.Compare(file,info) == 0 {
				fmt.Printf("Image unique: %s\n", text)
				fmt.Println(filepath.Abs(text))
			}
		}
	}
}