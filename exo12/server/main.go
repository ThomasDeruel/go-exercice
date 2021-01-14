package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
	"crypto/sha256"
	"sync"
)

type typeURLs struct {
	Urls []string `json:"urls,omitempty"`
}
var wg sync.WaitGroup

func hashImg(path string, wg *sync.WaitGroup) [32]byte{
	resp, err := http.Get(path)
	if err != nil {
		log.Fatal(err)
	}
	content, er := ioutil.ReadAll(resp.Body)
	fmt.Println(path)
	if er != nil {
		log.Fatal(er)
	}
	defer wg.Done()
	return sha256.Sum256(content)
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

		ch:= make(chan [32]byte)
		//my_bytes := [][]byte{}
		for _,url := range urls.Urls {
			wg.Add(1)
			go func(){
				ch <-hashImg(url, &wg)
			}()
			wg.Wait()
		}

		for res := range ch{
			fmt.Printf("Mon hash: %s \n", res)
		 }
	})
	http.ListenAndServe(":8081", nil)
}
