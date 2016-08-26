package main

import "testing"

var tests = []struct {
	n        int
	expected uint64
}{
	{10, 2520},
	{20, 232792560},
}

func TestFactorize(t *testing.T) {
	for _, tt := range tests {
		v := evenlyDivisible(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
