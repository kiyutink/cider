package main

import (
	"errors"
	"strconv"
	"strings"
)

func uint32ToIP(n uint32) string {
	mask := leastSignificant(8)

	stringOctets := []string{}

	for i := 3; i >= 0; i-- {
		octet := n >> (i * 8) & mask
		stringOctet := strconv.Itoa(int(octet))
		stringOctets = append(stringOctets, stringOctet)
	}

	return strings.Join(stringOctets, ".")
}

func ipToUint32(ip string) (uint32, error) {
	stringOctets := strings.Split(ip, ".")

	if len(stringOctets) != 4 {
		return 0, errors.New("invalid ip address")
	}

	octets := []uint32{}

	for _, stringOctet := range stringOctets {
		octet, err := strconv.Atoi(stringOctet)
		if err != nil {
			return 0, errors.New("invalid ip address")
		}
		octets = append(octets, uint32(octet))
	}

	n := uint32(0)
	for i, o := range octets {
		n += o << (8 * (3 - i))
	}

	return n, nil
}
