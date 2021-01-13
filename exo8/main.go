package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
)

type User struct {
	userName string
	Password string
}

func main() {
	data, err := ioutil.ReadFile("./users.json")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", users)
	/*
		output:

		[{userName: Password:123456} {userName: Password:azerty}]

	*/
}
