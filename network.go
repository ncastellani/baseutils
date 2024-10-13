package baseutils

import "net"

// GetFQDNIPv4Address
// fetch the first IPv4 address from an FQDN, ignoring errors
func GetFQDNIPv4Address(fqdn string) (addr string) {
	ips, _ := net.LookupIP(fqdn)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			addr = ipv4.String()
		}
	}

	return
}

// GetIPAddressRDNS
// fetch the reverse DNS address from an IP, ignoring errors
func GetIPAddressRDNS(addr string) (rdns string) {
	addresses, _ := net.LookupAddr(addr)
	return addresses[0]
}
