package main

import "github.com/bbetov/euler/shared/primer"

func matchesConjecture(n uint64, plist []uint64) bool {
	for _, p := range plist {
		if p > n-2 {
			break
		}
		rem := n - p
		if rem%2 != 0 {
			continue
		}
		rem /= 2
		for i := uint64(1); ; i++ {
			isq := i * i
			if isq == rem {
				return true
			}
			if isq > rem {
				break
			}
		}
	}
	return false
}

func main() {
	primes := make(map[uint64]bool)
	plist := []uint64{}
	var lastPrime uint64
	pr := primer.NewPrimer()

	for i := uint64(9); ; i = i + 2 {
		for lastPrime <= i {
			lastPrime, _ = pr.NextPrime()
			primes[lastPrime] = true
			plist = append(plist, lastPrime)
		}
		if _, ok := primes[i]; ok {
			continue // i is a prime number
		}
		if !matchesConjecture(i, plist) {
			println(i)
			break
		}
	}
}
