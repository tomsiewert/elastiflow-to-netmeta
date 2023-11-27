package misc

import "github.com/seancfoley/ipaddress-go/ipaddr"

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func ConvertIPv4ToIPv6(input string) string {
	address := ipaddr.NewIPAddressString(input).GetAddress()
	if address.IsIPv4() {
		convertedIPv4, _ := address.ToIPv4().GetIPv4MappedAddress()
		lastIp, _ := convertedIPv4.ToMixedString()
		return lastIp
	} else {
		return input
	}
}
