package main

import "testing"

var tests = []struct {
	n     []int
	total int
	e     int
}{
	{[]int{}, 200, 73682},
}

func TestWaysToSum(t *testing.T) {
	var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}
	for _, tt := range tests {
		v := waysToSum(tt.total, len(coins)-1, coins)
		if v != tt.e {
			t.Errorf("Expected %v, got %v", tt.e, v)
		}
	}
}

func BenchmarkClever(b *testing.B) {
	var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}
	for i := 0; i < b.N; i++ {
		waysToSum(200, len(coins)-1, coins)
	}
}

func BenchmarkBrute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waysToSumBruteForce(200)
	}
}
