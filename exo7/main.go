package main

import (
	"fmt"
)

func PrintIt(input interface{}) {
	// Or switch elem := ... => to get value
	switch input.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("string")
	case byte:
		fmt.Println("byte")
	default:
		fmt.Println(`Idk ¯\_(ツ)_/¯`)
	}
}
func main() {
	PrintIt("1")     // string
	PrintIt(1)       // int
	PrintIt(byte(1)) // byte
	PrintIt(1.0)     // float -> not defined
}
