package main

import "testing"

func TestSum(t *testing.T) {
	v := sum(10)
	if v != 23 {
		t.Error("Expected 23, got ", v)
	}
}

func TestSumEuler(t *testing.T) {
	v := sum(1000)
	if v != 233168 {
		t.Error("Expected 233168, got ", v)
	}
}

func BenchmarkSum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum(1000)
	}
}
