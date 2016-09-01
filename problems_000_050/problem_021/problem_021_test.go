package main

import "testing"

func TestLargeSumLast10Digits(t *testing.T) {
	v := getAmicableSum()
	if v != 31626 {
		t.Error("Expected ", 31626, " got ", v)
	}
}
