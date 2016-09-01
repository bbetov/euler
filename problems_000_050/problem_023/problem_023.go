package main

import (
	"fmt"
	"sort"

	"github.com/bbetov/euler/shared"
)

func getAbundantNumbers(limit int) []int {
	primes := shared.GetPrimesUInt64(uint64(limit))
	var ab []int
	for i := 12; i <= limit; i++ {
		s := shared.GetDivisorsSumPrimesSet(uint64(i), primes)
		if int(s) > i {
			ab = append(ab, i)
		}
	}
	return ab
}

func sumIntsNoAbSum() uint64 {
	limit := 28123
	ab := getAbundantNumbers(limit)
	sort.Ints(ab)
	as := make(map[int]bool)
	for _, i := range ab {
		for _, j := range ab {
			if i >= j && i+j <= limit {
				as[i+j] = true
			}
		}
	}
	var ts uint64
	for i := 1; i <= limit; i++ {
		if _, ok := as[i]; !ok {
			ts += uint64(i)
		}
	}
	return ts
}

func main() {
	fmt.Printf("%v", sumIntsNoAbSum())
}
