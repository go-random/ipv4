package ipv4

import (
	"fmt"
	"math/rand"
	"net"
)

func (nr *Randomizer) GenerateRandomIpv4() net.IP {
	ip := make(net.IP, 4)
	return ip.To4()
}

func (nr *Randomizer) GenerateRandomIpv4FromCidr(cidrRange string) (net.IP, error) {
	return GenerateRandomIpv4FromCidr(cidrRange)
}

func GenerateRandomIpv4FromCidr(cidrRange string) (net.IP, error) {
	_, network, err := net.ParseCIDR(cidrRange)
	if err != nil {
		return nil, fmt.Errorf("error parsing CIDR range: %w", err)
	}

	netmaskBytes := network.Mask
	if len(netmaskBytes) != net.IPv4len {
		return nil, fmt.Errorf("invalid netmask: must be either IPv4")
	}

	var addressBytes = network.IP.To4()

	if len(netmaskBytes) != len(addressBytes) {
		return nil, fmt.Errorf("netmask byte length does not match IP address byte length")
	}

	var picked []byte
	for i := 0; i < len(netmaskBytes); i++ {
		picked = append(picked, ((255^netmaskBytes[i])&byte(rand.Intn(256)))|addressBytes[i])
	}

	pickedIP := net.IP(picked)

	return pickedIP, nil
}
