package main

import (
	"fmt"
	"sort"

	"github.com/bbetov/euler/shared"
)

func main() {
	abundant := make(map[int]bool)
	var ab []int
	for i := 12; i <= 28123; i++ {
		s := shared.GetDivisorsSum(uint64(i))
		if int(s) > i {
			abundant[i] = true
			ab = append(ab, i)
		}
	}
	sort.Ints(ab)
	as := make(map[int]bool)
	for _, i := range ab {
		for _, j := range ab {
			if i >= j && i+j <= 28123 {
				as[i+j] = true
			}
		}
	}
	var ts uint64
	for i := 1; i <= 28123; i++ {
		if _, ok := as[i]; !ok {
			ts += uint64(i)
		}
	}
	fmt.Printf("%v", ts)
}
