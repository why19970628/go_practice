package main

import (
	"fmt"
	"net"
	"strings"
)

func validIPAddress(queryIP string) string {
	var ip = net.ParseIP(queryIP)
	if ip == nil {
		return "Neither"
	}
	if ip.To4() != nil {
		for _, s := range strings.Split(queryIP, ".") {
			if len(s) > 1 && s[0] == '0' {
				return "Neither"
			}
		}
		return "IPv4"
	}
	for _, s := range strings.Split(queryIP, ":") {
		if len(s) > 4 || s == "" {
			return "Neither"
		}
	}
	return "IPv6"
}

func main() {
	ip := "172.16.254.1"
	fmt.Println(validIPAddress(ip))
}
