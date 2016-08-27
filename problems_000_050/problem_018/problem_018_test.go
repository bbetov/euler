package main

import "testing"

func TestMaxPath(t *testing.T) {
	rn, _ := loadData("input.txt")
	v := rn.maxPath()
	if v != 1074 {
		t.Error("Expected ", 1074, " got ", v)
	}
}
