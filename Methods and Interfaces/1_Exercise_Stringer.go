package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (host IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v\n", host[0], host[1], host[2], host[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
		fmt.Printf(ip.IPPrint())
	}
}
