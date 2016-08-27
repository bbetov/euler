package main

import "testing"

func TestSundaysOnFirstOfTheMonth(t *testing.T) {
	v := sundaysOnFirst()
	if v != 171 {
		t.Error("Expected ", 171, " got ", v)
	}
}
