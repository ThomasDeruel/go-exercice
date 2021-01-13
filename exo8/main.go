package main

import (
	//"fmt"
	"encoding/json"
	"log"
	"os"
)

type User struct {
	Login    string `json:"user-Name"`
	Password string
}

func main() {
	var Paul = User{"Paul", "pass123"}
	paul_injson, err := json.Marshal(Paul)
	if err != nil {
		log.Fatal(err)
	}
	// stdout: sortie standard ->
	os.Stdout.Write(paul_injson)
}
