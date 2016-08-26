package main

import "testing"

var tests = []struct {
	n        int
	expected uint64
}{
	{5, 28},
	{500, 76576500},
}

func TestTriangularNumDivisors(t *testing.T) {
	for _, tt := range tests {
		v, _ := getTriangularNumber(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
