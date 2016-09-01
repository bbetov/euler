package main

import "testing"

func TestSumIntsNoAbSum(t *testing.T) {
	v := sumIntsNoAbSum()
	if v != 4179871 {
		t.Error("Expected ", 4179871, " got ", v)
	}
}

func TestAbundantNumbers(t *testing.T) {
	v := getAbundantNumbers(12)
	if len(v) != 1 || v[0] != 12 {
		t.Error("Expected single element slice[12]")
	}
}
