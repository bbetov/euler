package main

import "github.com/bbetov/euler/shared"

func getPentagonNumber(i uint64) (n uint64) {
	return (i * (i*3 - 1)) / 2
}

func isPentagon(P uint64) bool {
	// Compute n
	n := shared.SqrtUInt64(P*24 + 1)
	n = (n + 1) / 6
	return P == getPentagonNumber(n)
}

func main() {
	found := false
	diff := uint64(0)
	for i := uint64(2); ; i++ {
		j := getPentagonNumber(i)
		for q := i - 1; q >= 1; q-- {
			k := getPentagonNumber(q)
			d := j - k
			bd := isPentagon(d)
			s := j + k
			bs := isPentagon(s)
			if bd && bs {
				diff = d
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	println(diff)
}
