package main

import (
	"fmt"
	"sort"

	"github.com/bbetov/euler/shared"
)

func createHash(n uint64) (p uint64) {
	a := shared.Num2Array(n)
	sort.Sort(a)
	return shared.Array2Num(a)
}

func solution62(howManyCubes int) uint64 {
	processed := make(map[uint64]int)
	smallest := make(map[uint64]uint64)
	for n := uint64(345); ; n++ { // Until we find a match
		nCubed := n * n * n
		h := createHash(nCubed)
		processed[h]++
		v := processed[h]
		if _, ok := smallest[h]; !ok {
			smallest[h] = nCubed
		}
		if v == howManyCubes {
			return smallest[h]
		}
	}
}

func main() {
	fmt.Printf("%v\n", solution62(5))
}
