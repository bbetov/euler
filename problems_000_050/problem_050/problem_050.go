package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

func sumSlice(sl []uint64) (s uint64) {
	for _, p := range sl {
		s += p
	}
	return s
}

func main() {
	primes := shared.GetPrimesUInt64(uint64(1000000))
	primeSet := make(map[uint64]bool)
	psum := []uint64{0}
	s := uint64(0)
	for _, p := range primes {
		primeSet[p] = true
		s += p
		psum = append(psum, s)
	}

	maxLen := 0
	maxPrimeSum := uint64(0)
	lenPrimes := len(primes)
	start := 0
	for l := 2; l < lenPrimes-1; l++ {
		for st := 0; st < lenPrimes-1-l; st++ {
			s := psum[st+l] - psum[st]
			if s > 1000000 {
				break
			}
			if _, ok := primeSet[s]; ok {
				maxLen = l
				maxPrimeSum = s
				start = st
			}
		}
	}

	fmt.Printf("MaxPrime: %v, NumPrimes: %v, Start: %v", maxPrimeSum, maxLen, start)
}
