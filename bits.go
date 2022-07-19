package main

func leastSignificant(n int) uint32 {
	return ^uint32(0) << (32 - n) >> (32 - n)
}

func mostSignificant(n int) uint32 {
	return ^uint32(0) >> (32 - n) << (32 - n)
}
