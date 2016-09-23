package main

import "testing"

func TestSquareRootConvergents(t *testing.T) {
	v := getSquareRootConvergents(1000)
	if v != 153 {
		t.Error("SquareRootConvergents should be ", 153)
	}

	v = getSquareRootConvergents(9)
	if v != 1 {
		t.Error("SquareRootConvergents should be ", 1)
	}
}
