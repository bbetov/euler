package main

import "testing"

func TestPitagoreanTriples(t *testing.T) {
	tr := pitTripleGen()
	v := tr.a * tr.b * tr.c
	if v != 31875000 {
		t.Error("Expected ", 31875000, " got ", v)
	}
}
