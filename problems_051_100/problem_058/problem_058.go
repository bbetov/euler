package main

import "github.com/bbetov/euler/shared/primer"

func getPrimeChecker() func(uint64) bool {
	var pr = primer.NewPrimerWithSegSize(100000000)
	var np uint64

	return func(n uint64) bool {
		for n > np {
			np, _ = pr.NextPrime()
		}
		if n == np {
			return true
		}
		return false
	}
}
func findSideLen() uint64 {
	n := uint64(1)
	step := uint64(2)
	pc := getPrimeChecker()
	var np, nnp int

	for {
		for i := 0; i < 4; i++ {
			n += step

			if pc(n) {
				np++
			} else {
				nnp++
			}
			if np*10 < (np + nnp + 1) {
				return step + 1
			}
		}
		step += 2
	}
}

func main() {
	println(findSideLen())
}
