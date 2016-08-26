package main

import "testing"

func TestMaxProductFrom2DArray(t *testing.T) {
	v := getMaxProduct()
	if v != 70600674 {
		t.Error("Expected ", 70600674, " got ", v)
	}
}
