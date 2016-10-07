package main

import "testing"

func TestSquareRootConvergents(t *testing.T) {
	v := solution62(3)
	if v != 41063625 {
		t.Error("Smallest of 3 special square numbers should be ", 41063625)
	}

	v = solution62(5)
	if v != 127035954683 {
		t.Error("Smallest of 5 special square numbers should be ", 127035954683)
	}
}
