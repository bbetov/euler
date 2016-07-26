package main

import (
	"fmt"
	"math/big"

	"github.com/bbetov/euler/shared"
)

func div(dvsr, nd int) (r []byte) {
	cur := 1
	nz := 0
	for it := 0; cur > 0 && it < nd; it++ {
		if cur < dvsr {
			r = append(r, 0)
		} else {
			d := cur / dvsr
			r = append(r, byte(d))
			cur = cur % dvsr
			if nz == 0 {
				nz = len(r) - 1
				cur *= 10
				continue
			}
		}
		cur *= 10
	}
	return r
}

// http://mathworld.wolfram.com/DecimalExpansion.html
// Several options:
// 1. Finite decimal expansion ==> Numbers in the form: 2^a * 5^b
// 2. Repetitive Expansion ==> Numbers coprime to 10 (i.e. has no common divisors to with 10 except 1 - should not be divisable by both 2 and 5)

func compareSlices(a, b []byte) (equal bool) {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	maxNum := 1000
	maxSeq, numForMaxSeq := 0, 0
	for i := 3; i <= maxNum; i++ {
		m := shared.GetDivisorsFreq(uint64(i))
		_, has2 := m[uint64(2)]
		_, has5 := m[uint64(5)]

		if (has2 && has5 && len(m) == 2) || (len(m) == 1 && (has2 || has5)) {
			// Finite decimal expansion
			continue
		}

		if !has2 && !has5 { // Decimal expansion starts at the decimal point
			// (10 ^ k - 1) % q == 0, period k, divisor q
			q := big.NewInt(int64(i))
			zero := big.NewInt(0)
			ten := big.NewInt(10)
			for k := 1; k < i; k++ {
				z := big.NewInt(10)
				for ppp := 1; ppp < k; ppp++ {
					z.Mul(z, ten)
				}
				z.Sub(z, big.NewInt(1))
				z.Mod(z, q)
				if z.Cmp(zero) == 0 {
					fmt.Printf("%v ==> %v\n", i, k)
					if maxSeq < k {
						maxSeq = k
						numForMaxSeq = i
					}
					break
				}
			}
		}

	}
	fmt.Printf("Num with max seq: %v. Seq Len: %v", numForMaxSeq, maxSeq)

}
