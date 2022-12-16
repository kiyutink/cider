package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	list := flag.Bool("l", false, "list all the ip addresses instead of only start and end of the block")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, errors.New("invalid CIDR"))
		os.Exit(1)
	}

	cidr := flag.Args()[0]

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

	fmt.Printf("Total IP addresses: %v\n", 1 << (32 - maskInt))

	if *list {
		for addr := start; addr <= end; addr++ {
			fmt.Println(uint32ToIP(addr))
			if addr == ^uint32(0) {
				break
			}
		}
	} else {
		fmt.Println(uint32ToIP(start))
		fmt.Println(uint32ToIP(end))
	}
}
