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
	err = iprange.Add("153.168.3.36")
	if err != nil {
		panic(err)
	}
	err = iprange.Add("153.168.3.33")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", iprange.Stats)

	err = iprange.Delete("153.168.3.33")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", iprange.Stats)

	ips, err := ipranger.Ips("153.168.3.36/16")
	if err != nil {
		panic(err)
	}
	fmt.Println(len(ips))
}
