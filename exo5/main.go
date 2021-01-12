package main

import (
	"fmt"
	//"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	res := []string{}
	for _, val := range ip {
		// res = append(res,fmt.Sprintf((string(val)))) (same)
		res = append(res, fmt.Sprintf("%v", val))
	}
	stringval := strings.Join(res, ".")
	return fmt.Sprintf("print: %v", stringval)
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for _, val := range hosts {
		fmt.Println(val)
	}
}
