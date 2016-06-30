package main

import "testing"

func TestSum(t *testing.T) {
	v := sum(9)
	if v != 23 {
		t.Fatalf("Expected 23 and got %d", v)
	}
}
