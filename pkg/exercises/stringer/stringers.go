package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func String(ip IPAddr) string {
	return strings.ReplaceAll(fmt.Sprintf("%v", ip), " ", ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, String(ip))
	}
}
