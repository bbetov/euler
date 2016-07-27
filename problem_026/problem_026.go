package main

import (
	"fmt"
	"math/big"

	"github.com/bbetov/euler/shared"
)

// http://mathworld.wolfram.com/DecimalExpansion.html
// Several options:
// 1. Finite decimal expansion ==> Numbers in the form: 2^a * 5^b
// 2. Repetitive Expansion ==> Numbers coprime to 10 (i.e. has no common divisors to with 10 except 1 - should not be divisable by both 2 and 5)

func main() {
	maxNum := 1000
	maxSeq, numForMaxSeq := 0, 0
	for i := 12; i <= maxNum; i++ {
		t := i
		m := shared.GetDivisorsFreq(uint64(t))
		times2, has2 := m[uint64(2)]
		times5, has5 := m[uint64(5)]

		if (has2 && has5 && len(m) == 2) || (len(m) == 1 && (has2 || has5)) {
			// Finite decimal expansion
			continue
		}

		// Reduce the number by removing its 2 and 5 divisors
		for c := 0; c < times5; c++ {
			t /= 5
		}
		for c := 0; c < times2; c++ {
			t /= 2
		}

		// (10 ^ k - 1) % q == 0, period k, divisor q
		q := big.NewInt(int64(t))
		zero := big.NewInt(0)
		ten := big.NewInt(10)
		for k := 1; k < t; k++ {
			z := big.NewInt(10)
			for ppp := 1; ppp < k; ppp++ {
				z.Mul(z, ten)
			}
			z.Sub(z, big.NewInt(1))
			z.Mod(z, q)
			if z.Cmp(zero) == 0 {
				// fmt.Printf("%v ==> %v\n", i, k)
				if maxSeq < k {
					maxSeq = k
					numForMaxSeq = i
				}
				break
			}
		}
	}
	fmt.Printf("Num with max seq: %v. Seq Len: %v", numForMaxSeq, maxSeq)
}
