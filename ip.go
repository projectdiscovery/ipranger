package ipranger

import (
	"net"

	"github.com/projectdiscovery/iputil"
	"github.com/projectdiscovery/mapcidr"
)

// Ips of a cidr
func Ips(cidr string) ([]string, error) {
	return mapcidr.IPAddresses(cidr)
}

func ToCidr(item string) *net.IPNet {
	if iputil.IsIPv4(item) {
		item += "/32"
	} else if iputil.IsIPv6(item) {
		item += "/128"
	}
	if iputil.IsCIDR(item) {
		_, ipnet, _ := net.ParseCIDR(item)
		return ipnet
	}

	return nil
}
