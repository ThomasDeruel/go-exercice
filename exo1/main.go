package main

import (
	"fmt"
	"flag"
)

func main() {
	const VERSION = "1.0"
	myversion := flag.Bool("version",false, "show the program version")
	flag.Parse()
	if *myversion {
		fmt.Println(VERSION)
		return
	}
	fmt.Println("Hello World")
	
}
