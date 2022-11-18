package main

import (
	"fmt"

	"github.com/projectdiscovery/ipranger"
)

func main() {
	iprange, err := ipranger.New()
	if err != nil {
		panic(err)
	}
	err = iprange.Add("127.0.0.1")
	if err != nil {
		panic(err)
	}
	err = iprange.Add("127.0.100.10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", iprange.Stats)

	err = iprange.Delete("127.0.100.10")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", iprange.Stats)

	ips, err := ipranger.Ips("127.0.0.1/16")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(ips))
}
