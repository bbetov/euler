package main

import "testing"

var tests = []struct {
	n        uint64
	expected uint64
}{
	{2000000, 142913828922},
	{10, 17},
}

func TestPrimesSum(t *testing.T) {
	for _, tt := range tests {
		v := getPrimesSum(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
