package medium

import (
	"net"
	"strings"
)

func validIPAddress(IP string) string {
	
	ipstr := net.ParseIP(IP)
	if ipstr == nil {
		return "Neither"
	} else if strings.Contains(IP,":") {
		return "IPv6"
	} else {
		return "IPv4"
	}
}