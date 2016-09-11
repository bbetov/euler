package main

import "github.com/bbetov/euler/shared/primer"

func nextBatchPrimes(pr *primer.Primer, lastPrime, rangeEnd uint64) (rv []uint64, lp, re uint64) {
	for {
		if lastPrime > 0 {
			rv = append(rv, lastPrime) // the last from the previous range
			lastPrime = 0
		}
		p, _ := (*pr).NextPrime()

		if p < rangeEnd {
			rv = append(rv, p)
			continue
		}
		lastPrime = p
		break
	}
	rangeEnd *= 10
	return rv, lastPrime, rangeEnd
}

func num2Array(p uint64) (rv []byte) {
	for p > 0 {
		r := p % 10
		p /= 10
		rv = append(rv, byte(r))
	}
	return rv
}

func altPotentialPrime(pa []byte, dig byte, mask int) (p uint64) {
	t := 1 << uint(len(pa))
	for i := len(pa) - 1; i >= 0; i-- {
		t >>= 1
		pr := pa[i]
		p *= 10
		if mask&t > 0 {
			pr = dig
		}
		p += uint64(pr)
	}

	return p
}

func findPattern(pl []uint64, fr int) uint64 {
	pmap := make(map[uint64]bool)
	for _, tmp := range pl {
		pmap[tmp] = true
	}

	for _, p := range pl {
		pa := num2Array(p)
		uv := (1 << uint(len(pa)))

		for m := 0; m < uv; m++ {
			count := 0
			min := uint64(0)
			for i := byte(0); i < 10; i++ {
				if i == 0 && m&(1<<uint(len(pa)-1)) > 0 {
					continue // skip zeroing the first digit
				}
				ap := altPotentialPrime(pa, i, m)
				if _, ok := pmap[ap]; ok {
					if min > ap || min == 0 {
						min = ap
					}
					count++
				}
			}
			if count == fr {
				return min
			}
		}
	}
	return 0
}

func findPrimeDigitReplacement(familyRank int) (p uint64) {
	pr := primer.NewPrimerWithSegSize(1000000)
	rangeEnd := uint64(100)
	lastPrime := uint64(0)
	_, lastPrime, rangeEnd = nextBatchPrimes(&pr, lastPrime, rangeEnd) // to initialize

	for {
		var pl []uint64
		pl, lastPrime, rangeEnd = nextBatchPrimes(&pr, lastPrime, rangeEnd)
		p = findPattern(pl, familyRank)
		if p > 0 {
			break
		}
	}
	return p
}

func main() {
	println(findPrimeDigitReplacement(7))
}
