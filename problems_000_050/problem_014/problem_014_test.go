package main

import "testing"

func TestMaxCollatzLen(t *testing.T) {
	_, v := getMaxCollatzSeq()
	if v != 837799 {
		t.Error("Expected ", 837799, " got ", v)
	}
}

func TestCollatzLen(t *testing.T) {
	v := collatzLen(int64(13))
	if v != 10 {
		t.Error("Expected ", 10, " got ", v)
	}
}
