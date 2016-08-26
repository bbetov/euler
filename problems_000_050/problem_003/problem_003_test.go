package main

import "testing"

var tests = []struct {
	n        uint64 // input
	expected uint64 // expected result
}{
	{42, 7},
	{830297, 17},
	{600851475143, 6857},
}

func TestGetLargestPrimeFactor(t *testing.T) {
	for _, tt := range tests {
		v := getLargestPrimeFactor(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}

func BenchmarkFibEuler(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getLargestPrimeFactor(600851475143)
	}
}
