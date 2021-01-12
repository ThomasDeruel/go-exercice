package main

import (
	"fmt"
	"time"
)

type MyError struct{
	When time.Time
	What string
}

func (err *MyError) Error() string {
	return fmt.Sprintf("[Error]:\nWhen %s: \nWhat: %s", err.When, err.What)
}

func run() error {
	return &MyError{
		When: time.Now(),
		What:"ERROR 404 WTF",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
