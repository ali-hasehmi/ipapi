package main

import (
	"log"

	"github.com/ali-hasehmi/ipapi"
)

func main() {
	ipv4, err := ipapi.QueryOwnIPv4()
	if err != nil {
		log.Fatalln("error finding ipv4:", err)
	}
	ipv6, err := ipapi.QueryOwnIPv6()
	if err != nil {
		log.Fatalln("error finding ipv4:", err)
	}
	if ipv4 == "" {
		ipv4 = "none"
	}
	if ipv6 == "" {
		ipv6 = "none"
	}
	log.Printf("your ipv4: %v\n", ipv4)
	log.Printf("your ipv6: %v\n", ipv6)
}
