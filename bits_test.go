package main

import (
	"testing"
)

func TestUint32ToIP(t *testing.T) {
	tt := []struct {
		expected string
		integer  uint32
	}{
		{
			"255.255.255.255",
			4294967295,
		},
		{
			"27.13.101.56",
			453862712,
		},
		{
			"0.0.0.0",
			0,
		},
	}

	for _, test := range tt {
		if uint32ToIP(test.integer) != test.expected {
			t.Errorf("expected %v to convert to ip %v, instead got %v",
				test.integer, test.expected, uint32ToIP(test.integer))
		}
	}
}

func TestIPToUint32(t *testing.T) {
	tt := []struct {
		ip       string
		expected uint32
		invalid  bool
	}{
		{
			"255.255.255.255",
			4294967295,
			false,
		},
		{
			"27.13.101.56",
			453862712,
			false,
		},
		{
			"0.0.0.0",
			0,
			false,
		},
		{
			"",
			0,
			true,
		},
		{
			"1.1.1",
			0,
			true,
		},
		{
			"a.b.c.d",
			0,
			true,
		},
	}

	for _, test := range tt {
		n, err := ipToUint32(test.ip)

		if err != nil && !test.invalid {
			t.Errorf("expected no error")
		}

		if err == nil && test.invalid {
			t.Errorf("expected an error")
		}

		if n != test.expected {
			t.Errorf("expected %v to convert to ip %v, instead got %v",
				test.ip, test.expected, n)
		}
	}
}
