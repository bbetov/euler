package main

import (
	"testing"

	"github.com/bbetov/euler/shared"
)

var tests = []struct {
	n        int
	expected uint64
}{
	{6, 13},
	{10001, 104743},
}

func TestGetPrimeUInt64(t *testing.T) {
	for _, tt := range tests {
		v := shared.GetPrimeUInt64(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
