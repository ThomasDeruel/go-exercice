package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte


// func(ip IPAddr) String() string {
// 	res := []string{}
// 	for _,val := range ip {
// 		res = append(res,strconv.Itoa(int(val)))
// 	}
// 	return fmt.Printf("print: %v", strings.Join(res,"."))
// }
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	// for _,val := range hosts {
	// 	fmt.Println(val)
	// }
	for _,ip := range hosts {
		res := []string{}
		for _,val := range ip {
			res = append(res,strconv.Itoa(int(val)))
		}
		fmt.Printf("print: %v\n", strings.Join(res,"."))
	}
}
