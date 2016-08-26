package main

import "testing"

var tests = []struct {
	n        int
	expected int
}{
	{4, 5832},
	{13, 23514624000},
}

func TestAdjacentProduct(t *testing.T) {
	for _, tt := range tests {
		v := adjacentProduct(tt.n)
		if v != tt.expected {
			t.Error("Expected ", tt.expected, " got ", v)
		}
	}
}
