package main

import (
	"fmt"
	"net/http"
	// "io/ioutil"
	"bytes"
	"log"
	"flag"
	"strings"
	//"os"
	"encoding/json"
	// "crypto/sha256"
)
type typeURLs struct {
	Urls []string `json:"urls,omitempty"`
}
func main(){
	var list_url typeURLs
	urls := flag.String("urls", "", "send list urls"); 
	flag.Parse()
	if *urls != "" {
		list_url = typeURLs{strings.Split(*urls,",")}
		urls_byte, err := json.Marshal(list_url)
		if err != nil {
			log.Fatal(err)
			return
		}
		reader := bytes.NewReader(urls_byte)
		// response, er => get response
		_, er := http.Post(
			"http://localhost:8081",
			"appplication/json",
			reader,
		)
		if er != nil {
			log.Fatal(er)
		}
		fmt.Println("ok")
	}
	fmt.Println("-urls is empty")
}