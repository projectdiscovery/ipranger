package ipranger

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/projectdiscovery/mapcidr"
)

type Item struct {
	IP   string
	Port int
}

func (i Item) String() string {
	return net.JoinHostPort(i.IP, fmt.Sprintf("%d", i.Port))
}

func randint64() int64 {
	var b [8]byte
	rand.Read(b[:])
	return int64(binary.LittleEndian.Uint64(b[:]))
}

type IPShuffle struct {
	ipranger *IPRanger
	Ports    []int
}

func NewIPShuffle() (*IPShuffle, error) {
	ipr, err := New()
	if err != nil {
		return nil, err
	}
	return &IPShuffle{ipranger: ipr}, nil
}

func (i *IPShuffle) Shuffle() (chan Item, error) {
	return i.ShuffleWithSeed(randint64()), nil
}

func (i *IPShuffle) ShuffleWithSeed(seed int64) chan Item {
	out := make(chan Item)
	go func(out chan Item) {
		defer close(out)
		targetsCount := int64(i.ipranger.Stats.IPS)
		portsCount := int64(len(i.Ports))
		Range := targetsCount * portsCount
		br := NewBlackRock(Range, seed)
		for index := int64(0); index < Range; index++ {
			xxx := br.Shuffle(index)
			ipIndex := xxx / portsCount
			portIndex := int(xxx % portsCount)
			ip := i.PickIP(ipIndex)
			port := i.PickPort(portIndex)

			if ip == "" || port <= 0 {
				continue
			}
			out <- Item{IP: ip, Port: port}
		}
	}(out)
	return out
}

func (i *IPShuffle) PickIP(index int64) string {
	for _, target := range i.ipranger.CoalescedHostList {
		subnetIpsCount := int64(mapcidr.AddressCountIpnet(target))
		if index < subnetIpsCount {
			return i.PickSubnetIP(target, index)
		}
		index -= subnetIpsCount
	}

	return ""
}

func (i *IPShuffle) PickSubnetIP(network *net.IPNet, index int64) string {
	return mapcidr.Inet_ntoa(mapcidr.Inet_aton(network.IP) + index).String()
}

func (i *IPShuffle) PickPort(index int) int {
	return i.Ports[index]
}
