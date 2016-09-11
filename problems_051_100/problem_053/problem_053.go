package main

import "fmt"

type bcKey struct {
	n int
	k int
}

var bcMap = make(map[bcKey]uint64)

func calculateBinomialCoeff(n, k int) (rv uint64) {
	if k == 0 || k == n {
		return 1
	}
	var bc1, bc2 uint64
	var ok bool

	k1 := bcKey{n: n - 1, k: k - 1}
	if bc1, ok = bcMap[k1]; !ok {
		bc1 = calculateBinomialCoeff(k1.n, k1.k)
		bcMap[k1] = bc1
	}

	k2 := bcKey{n: n - 1, k: k}
	if bc2, ok = bcMap[k2]; !ok {
		bc2 = calculateBinomialCoeff(k2.n, k2.k)
		bcMap[k2] = bc2
	}

	return bc1 + bc2
}

func main() {
	var counter int
	for n := 1; n <= 100; n++ {
		for k := 1; k < n; k++ {
			if calculateBinomialCoeff(n, k) > 1000000 {
				fmt.Printf("%d, %d\n", n, k)
				counter++
			}
		}
	}
	fmt.Printf("\n\nAnswer: %d\n", counter)
}
