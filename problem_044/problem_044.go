package main

import "github.com/bbetov/euler/shared"

func getPentagonalNumber(i uint64) (n uint64) {
	return (i * (i*3 - 1)) / 2
}

func isPentagonal(P uint64) bool {
	// Compute n
	n := shared.SqrtUInt64(P*24 + 1)
	n = (n + 1) / 6
	return P == getPentagonalNumber(n)
}

func main() {
	found := false
	diff := uint64(0)
	for i := uint64(2); ; i++ {
		j := getPentagonalNumber(i)
		for q := i - 1; q >= 1; q-- {
			k := getPentagonalNumber(q)
			d := j - k
			bd := isPentagonal(d)
			s := j + k
			bs := isPentagonal(s)
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
