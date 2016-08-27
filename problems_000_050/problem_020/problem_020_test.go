package main

import "testing"

var tests = []struct {
	n        int
	expected uint64
}{
	{10, 27},
	{100, 648},
}

func TestSumDigitsFact(t *testing.T) {
	for _, tt := range tests {
		v := getSumDigitsFact(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
