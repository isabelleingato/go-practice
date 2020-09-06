package main

import(
	"fmt"
	"strings"
	"strconv"
);

type IPAddr [4]byte

func (ipaddr IPAddr) String() string {
	sBuffer := make([]string, len(ipaddr))
	for i, value := range ipaddr {
		sBuffer[i] = strconv.Itoa(int(value))
	}
	return strings.Join(sBuffer, ".") 
}

func ipaddr() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}