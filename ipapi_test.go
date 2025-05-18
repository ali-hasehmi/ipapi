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
		t.Fatal("received string as ipv4 is not ipv4: ", ip)
	}
	t.Logf("your machine's ipv6: %v", ip)
}
