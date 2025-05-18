package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ali-hasehmi/ipapi"
)

func main() {
	info, err := ipapi.QueryIPInfo("1.1.1.1")
	if err != nil {
		log.Fatalln(info)
	}
	b, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		log.Fatalln(info)
	}
	fmt.Println(string(b))
}
