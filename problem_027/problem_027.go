package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

var primes = make(map[uint64]bool)

func consecutivePrimes(a, b int64) (np int) {
	n := int64(0)
	for {
		res := n*n + a*n + b
		if res < 0 {
			break
		}
		if _, ok := primes[uint64(res)]; !ok {
			break
		}
		np++
		n++
	}

	return np
}

// n*n + a*n + b
// b must be positive, so [0 - 1000)
// a (-1000, 1000)
// n [0, ?]
func main() {
	p := shared.GetPrimesInt64(uint64(10000000))
	for _, tp := range p {
		primes[tp] = false
	}
	var maxPrimes int
	var maxa, maxb int64
	for a := int64(-999); a < 999; a++ {
		for b := int64(0); b < 999; b++ {
			numPrimes := consecutivePrimes(a, b)
			if numPrimes > maxPrimes {
				maxPrimes = numPrimes
				maxa = a
				maxb = b
				fmt.Printf("MaxPrimes: %v for a, b = %v, %v\n", maxPrimes, a, b)
			}
		}
	}
	fmt.Printf("MaxPrimes overall: %v for a, b = %v, %v\n", maxPrimes, maxa, maxb)
	fmt.Printf("Product is: %v", maxa*maxb)
}
