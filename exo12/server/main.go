package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"crypto/sha256"
)

type typeURLs struct {
	Urls []string `json:"urls,omitempty"`
}

func hashImg(path string) []byte{
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	content, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		log.Fatal(er)
	}
	h := sha256.New()
	h.Write(content)
	return h.Sum(nil)
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Errorreadingbody:%v", err)
			http.Error(
				rw,
				"can'treadbody", http.StatusBadRequest,
			)
			return
		}
		var urls typeURLs
		err = json.Unmarshal(body,&urls)
		if err != nil {
			log.Fatal(err)
			return
		}
		c := make(chan []byte)
		my_bytes := [][]byte{}
		for _,url := range urls.Urls {
			go func(){
				fmt.Println(url)
				//test := hashImg(url)
				c <- hashImg(url)
				my_bytes = append(my_bytes,<-c)
			}()
			fmt.Println(my_bytes)
		}
	})
	http.ListenAndServe(":80", nil)
}
