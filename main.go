package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, errors.New("invalid CIDR"))
		os.Exit(1)
	}

	cidr := os.Args[1]

	parts := strings.Split(cidr, "/")

	if len(parts) != 2 {
		fmt.Fprintln(os.Stderr, errors.New("invalid CIDR"))
		os.Exit(1)
	}

	ipStr := parts[0]

	addr, err := ipToUint32(ipStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.New("invalid CIDR IP"))
		os.Exit(1)
	}

	maskStr := parts[1]
	maskInt, err := strconv.Atoi(maskStr)
	if err != nil || maskInt < 0 || maskInt > 32 {
		fmt.Fprintln(os.Stderr, errors.New("invalid CIDR mask"))
		os.Exit(1)
	}

	mask := mostSignificant(maskInt)
	start := addr & mask
	end := start + ^mask

	fmt.Println(uint32ToIP(start), uint32ToIP(end))
}
