package ipranger

import (
	"errors"
	"github.com/projectdiscovery/mapcidr"
)

// IPs returns the list of IP addresses within a CIDR range.
func IPs(cidr string) ([]string, error) {
	ips, err := mapcidr.IPAddresses(cidr)
	if err != nil {
		return nil, errors.New("failed to retrieve IP addresses: " + err.Error())
	}
	return ips, nil
}
