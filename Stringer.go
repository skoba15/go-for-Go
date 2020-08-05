package main

import (
	"fmt"
)

type IPAddr [4]byte

// fmt.Printf() looks for String() string implementation for the given type, here we are overriding one
func (ip IPAddr) String() string {
	return  fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}