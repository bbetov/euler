package main

import "testing"

func TestMaxSumDigits(t *testing.T) {
	v := getMaxSumDigits()
	if v != 972 {
		t.Error("MaxSumDigits should be ", 972)
	}
}
