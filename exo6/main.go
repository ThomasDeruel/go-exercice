package main

import (
	"fmt"
	"time"
)

type MyError struct{
	When time.Time
	What string
}
// implementation, but ....
// Wtf is " * "
func (err *MyError) Error() string {
	return fmt.Sprintf("When %s: \nWhat: %s", err.When, err.What)
}
// Wtf is " & "
func run() error {
	return &MyError{
		When: time.Now(),
		What:"ERROR 404 WTF",
	}
}

func main() {
	fmt.Println(run())
}
