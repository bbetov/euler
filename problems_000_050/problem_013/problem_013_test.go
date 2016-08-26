package main

import "testing"

func TestLargeSumLast10Digits(t *testing.T) {
	v := getLast10LargeSum()
	if v != "5537376230" {
		t.Error("Expected ", "5537376230", " got ", v)
	}
}
