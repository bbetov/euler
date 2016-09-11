package main

import "testing"

var tests = []struct {
	n uint64
	e uint64
}{
	{2, 125874},
	{6, 142857},
}

func TestPermutedMultiples(t *testing.T) {
	for _, tt := range tests {
		v := findPermutedMult(tt.n)
		if v != tt.e {
			t.Error("Expected ", tt.e, " got ", v)
		}
	}
}
