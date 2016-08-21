package main

import (
	"github.com/bbetov/euler/shared"
	"github.com/bbetov/euler/shared/primer"
)

func main() {
	primes := make(map[uint64]bool)
	plist := []uint64{}
	var lastPrime uint64
	pr := primer.NewPrimer()
	for i := uint64(210); ; i++ {
		for lastPrime <= i+3 {
			lastPrime, _ = pr.NextPrime()
			primes[lastPrime] = true
			plist = append(plist, lastPrime)
		}
		oneIsPrime := false
		for pf := uint64(0); pf < 4; pf++ {
			if _, ok := primes[i]; ok {
				oneIsPrime = true
				break
			}
		}
		if oneIsPrime {
			continue
		}

		found := true
		for pf := uint64(0); pf < 4; pf++ {
			dvs := shared.GetDivisorsFreqPrimesSet(i+pf, plist)
			if len(dvs) != 4 {
				found = false
				break
			}
		}

		if !found {
			continue
		}
		println(i)
		break
	}
}
