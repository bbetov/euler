package main

import "testing"

func TestFibEuler(t *testing.T) {
	v := fibEven(4000000)
	if v != 4613732 {
		t.Error("Expected 4613732, got ", v)
	}
}

func BenchmarkFibEuler(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fibEven(4000000)
	}
}
