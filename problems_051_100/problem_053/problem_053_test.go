package main

import "testing"

var tests = []struct {
	n, k int
	e    uint64
}{
	{5, 3, 10},
	{29, 14, 77558760},
}

func TestBinomialCoeff(t *testing.T) {
	for _, tt := range tests {
		v := calculateBinomialCoeff(tt.n, tt.k)
		if v != tt.e {
			t.Error("Expected ", tt.e, " got ", v)
		}
	}
}
