package main

import "testing"

var tests = []struct {
	n        int
	expected uint64
}{
	{10, 2640},
	{100, 25164150},
}

func TestDiff(t *testing.T) {
	for _, tt := range tests {
		v := getDiff(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
