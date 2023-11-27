package misc

import "net/netip"

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ConvertIPv4ToIPv6(input string) string {
	if address, _ := netip.ParseAddr(input); address.Is4() {
		return netip.AddrFrom16(address.As16()).String()
	}
	return input
}
