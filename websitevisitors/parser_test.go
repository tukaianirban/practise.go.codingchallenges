package websitevisitors

import "testing"

func TestGetNPopularVisitors(t *testing.T) {

	ipaddrs := GetNPopularVisitors(3, "./samplelogs.txt")
	if len(ipaddrs) == 0 {
		t.Fatalf("found no ip addresses")
	}

	// t.Logf("found ip addresses in sorted order = %#v", ipaddrs)
}
