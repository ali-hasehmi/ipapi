package ipapi

import (
	"net/netip"
	"testing"
)

func TestQueryOwnIPv4(t *testing.T) {
	ip, err := QueryOwnIPv4()
	if err != nil {
		t.Fatal(err)
	}
	if ip == "" {
		t.Log("your machine has no public ipv4")
		return
	}
	nip, err := netip.ParseAddr(ip)
	if err != nil || !nip.Is4() {
		t.Fatal("received string as ipv4 is not ipv4: ", ip)
	}
	t.Logf("your machine's ipv4: %v", ip)
}

func TestQueryOwnIPv6(t *testing.T) {
	ip, err := QueryOwnIPv6()
	if err != nil {
		t.Fatal(err)
	}
	if ip == "" {
		t.Log("your machine has no public ipv6")
		return
	}
	nip, err := netip.ParseAddr(ip)
	if err != nil || !nip.Is6() {
		t.Fatal("received string as ipv6 is not ipv6: ", ip)
	}
	t.Logf("your machine's ipv6: %v", ip)
}

func TestQueryIpInfo(t *testing.T) {
	info, err := QueryIPInfo("1.1.1.1")
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal("no error reported but nil IpInfo returned")
	}
	if info.IPv4 != "1.1.1.1" {
		t.Fatal("fatal misinformation in returned IpInfo: IPv4 doesn't match")
	}
	if info.IPv6 != "" {
		t.Fatal("fatal misinformation in returned IpInfo: IPv6 doesn't match")
	}

	info, err = QueryIPInfo("2606:4700:4700::1111")
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal("no error reported but nil IpInfo returned")
	}
	if info.IPv4 != "" {
		t.Fatal("fatal misinformation in returned IpInfo: IPv4 doesn't match")
	}
	if info.IPv6 != "2606:4700:4700::1111" {
		t.Fatal("fatal misinformation in returned IpInfo: IPv6 doesn't match")
	}
}

func TestQueryOwnIPInfo(t *testing.T) {
	info, err := QueryOwnIPInfo()
	if err != nil {
		t.Fatal(err)
	}
	if info == nil {
		t.Fatal("no error reported but nil IpInfo returned")
	}
}
