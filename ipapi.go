package ipapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/netip"
	"time"
)

const (
	ipv4URL          = "https://api.ipify.org/"
	ipv6URL          = "https://api6.ipify.org/"
	queryURL         = "https://api.ipquery.io/"
	ipv4Connectivity = "8.8.8.8:53"
	ipv6Connectivity = "82001:4860:4860::888:53"
)

// ISPInfo represents information about the ISP of an IP address.
type ISPInfo struct {
	ASN string `json:"asn,omitempty"`
	Org string `json:"org,omitempty"`
	ISP string `json:"isp,omitempty"`
}

// LocationInfo represents geographical information about an IP address.
type LocationInfo struct {
	Country     string  `json:"country,omitempty"`
	CountryCode string  `json:"country_code,omitempty"`
	City        string  `json:"city,omitempty"`
	State       string  `json:"state,omitempty"`
	ZipCode     string  `json:"zipcode,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	Timezone    string  `json:"timezone,omitempty"`
	Localtime   string  `json:"localtime,omitempty"`
}

// RiskInfo represents risk information about an IP address.
type RiskInfo struct {
	IsMobile     bool `json:"is_mobile,omitempty"`
	IsVPN        bool `json:"is_vpn,omitempty"`
	IsTor        bool `json:"is_tor,omitempty"`
	IsProxy      bool `json:"is_proxy,omitempty"`
	IsDatacenter bool `json:"is_datacenter,omitempty"`
	RiskScore    int  `json:"risk_score,omitempty"`
}

// IPInfo represents all the information returned by the API.
type IPInfo struct {
	IPv4     string        `json:"ipv4"`
	IPv6     string        `json:"ipv6"`
	ISP      *ISPInfo      `json:"isp,omitempty"`
	Location *LocationInfo `json:"location,omitempty"`
	Risk     *RiskInfo     `json:"risk,omitempty"`
}

func checkConnectivityIPv4() bool {
	c, err := net.DialTimeout("tcp4", ipv4Connectivity, time.Duration(2*time.Second))
	if err != nil {
		return false
	}
	c.Close()
	return true
}

func checkConnectivityIPv6() bool {
	c, err := net.DialTimeout("tcp6", ipv6Connectivity, time.Duration(2*time.Second))
	if err != nil {
		return false
	}
	c.Close()
	return true
}

// Find this machine's public ipv4
// returns empty string if this machine doesn't have ipv4
func QueryOwnIPv4() (string, error) {
	resp, err := http.Get(ipv4URL)
	if err != nil {
		if checkConnectivityIPv4() {
			return "", err
		}
		if checkConnectivityIPv6() {
			return "", nil
		}
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch ip: %v", resp.Status)
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

// Find this machine's public ipv6
// returns empty string if this machine doesn't have ipv6
func QueryOwnIPv6() (string, error) {
	resp, err := http.Get(ipv6URL)
	if err != nil {
		if checkConnectivityIPv6() {
			return "", err
		}
		if checkConnectivityIPv4() {
			return "", nil
		}
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch ip: %v", resp.Status)
	}
	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

// Query a specific IP(v4 or v6) for details
func QueryIPInfo(ip string) (*IPInfo, error) {
	resp, err := http.Get(queryURL + ip)
	if err != nil {
		return nil, err
	}
	var ipinfo IPInfo
	nip, err := netip.ParseAddr(ip)
	if err != nil {
		return nil, err
	}
	if nip.Is4() {
		ipinfo.IPv4 = ip
	} else {
		ipinfo.IPv6 = ip
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ipinfo)
	if err != nil {
		return nil, err
	}
	return &ipinfo, nil
}

// Fetch Both IPv4 and IPv6 of this machine and their relevant details
func QueryOwnIPInfo() (*IPInfo, error) {
	ipv4, err := QueryOwnIPv4()
	if err != nil {
		return nil, err
	}
	ipv6, err := QueryOwnIPv6()
	if err != nil {
		return nil, err
	}

	var info *IPInfo
	if ipv4 != "" {
		info, err = QueryIPInfo(ipv4)
		if err != nil {
			return nil, err
		}
	} else {
		info, err = QueryIPInfo(ipv6)
		if err != nil {
			return nil, err
		}
	}
	info.IPv4 = ipv4
	info.IPv6 = ipv6

	return info, nil
}
