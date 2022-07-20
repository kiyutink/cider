package main

import (
	"testing"
)

func TestMostSignificant(t *testing.T) {
	tt := []struct {
		n        int
		expected uint32
	}{
		{0, 0},
		{1, 2147483648},
	}

	for _, test := range tt {
		if v := mostSignificant(test.n); v != test.expected {
			t.Errorf("expected %v most significant bits to be equal to %v, instead got %v", test.n, test.expected, v)
		}
	}
}

func TestLeastSignificant(t *testing.T) {
	tt := []struct {
		n        int
		expected uint32
	}{
		{0, 0},
		{1, 1},
		{4, 15},
		{10, 1023},
	}

	for _, test := range tt {
		if v := leastSignificant(test.n); v != test.expected {
			t.Errorf("expected %v least significant bits to be equal to %v, instead got %v", test.n, test.expected, v)
		}
	}
}
