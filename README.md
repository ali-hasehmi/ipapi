# ipapi

`ipapi` is a simple and lightweight Go package to fetch public IP address information (IPv4 and IPv6) of the current machine or any arbitrary IP address using third-party [APIs](###Acknoledgments)

## Features

- 🔍 Get your machine's public **IPv4** and **IPv6** addresses separately.
- 🌐 Lookup detailed geolocation and network info for any IP address.
- ⚡ Uses two fast APIs: `ipify.org` for IP fetching, and `ipquery.io` for metadata.
- 📦 Simple to use, easy to integrate.

## Installation

```bash
go get github.com/ali-hasehmi/ipapi
```

## Usage Examples

### **find my IPs**

```go
package main

import (
    "github.com/ali-hasehmi/ipapi"
)

func main(){
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
    if ipv6 == ""{
        ipv6= "none"
    }
    log.Printf("your ipv4: %v\n",ipv4)
    log.Printf("your ipv6: %v\n",ipv6)
}
```

### **Find Everything About my machine IPs**

```go
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
```

### **Find Everything About any IP**

```go
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
```

## Acknoledgments

- This package uses the [ipvifiy.org](https://www.ipify.org/) and [ipquery.io](https://ipquery.gitbook.io/ipquery-docs) APIs
- Inspired by Similiar Go package [ipapi-go](https://github.com/ipqwery/ipapi-go).
