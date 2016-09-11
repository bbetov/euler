package main

import "testing"

var tests = []struct {
	n int
	e uint64
}{
	{7, 56003},
	{8, 121313},
}

func TestPrimeDigitReplacement(t *testing.T) {
	for _, tt := range tests {
		v := findPrimeDigitReplacement(tt.n)
		if v != tt.e {
			t.Error("Expected ", tt.e, " got ", v)
		}
	}
}
