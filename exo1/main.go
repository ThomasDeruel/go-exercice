package main

import (
	"fmt"
	"flag"
)

func main() {
	const VERSION = "1.0"
	myversion := flag.Bool("version",true, "show the program version")
	flag.Parse()
	if *myversion {
		fmt.Println(myversion)
		return
	}
	fmt.Println("Hello World")
	
}
