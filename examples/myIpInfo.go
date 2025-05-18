package main

import (
	"encoding/json"
	"log"

	"github.com/ali-hasehmi/ipapi"
)

func main() {
	info, err := ipapi.QueryOwnIPInfo()
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(info, "", "\t")
	if err != nil {
		log.Fatalln(info)
	}
	log.Println(string(b))
}
