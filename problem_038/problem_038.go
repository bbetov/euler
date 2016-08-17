package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

func getNum(s []int) (n int) {
	mul := 1
	for i := len(s) - 1; i >= 0; i-- {
		n += s[i] * mul
		mul *= 10
	}
	return n
}

func fillBits(n int, fl uint16) (uint16, bool) {
	for n > 0 {
		r := uint(n % 10)
		bit := uint16(1 << r)
		if r == 0 || fl&bit > 0 { // already found or zero
			return fl, false
		}
		fl |= bit
		n /= 10
	}
	return fl, true
}

func processPermutation(perm []int) (mul int, ok bool) {
	for l := 1; l < len(perm)-1; l++ {
		n := getNum(perm[:l])
		if n == 1 {
			continue
		}
		ci := 0
		var br bool
		for i := 1; i < 9; i++ {
			if ci == 9 {
				break
			}
			p := n * i
			tmp := p
			var check []int
			for tmp > 0 {
				r := tmp % 10
				tmp /= 10
				check = append(check, r)
			}
			for q := len(check) - 1; q >= 0; q-- {
				if ci >= len(perm) || perm[ci] != check[q] {
					br = true
					break
				}
				ci++
			}
			if br {
				break
			}
		}
		if !br {
			return n, true
		}
	}
	return 0, false
}

func main() {
	perm := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	/* tests
	pp, oo := processPermutation([]int{1, 9, 2, 3, 8, 4, 5, 7, 6})
	fmt.Printf("%v, %v\n", pp, oo)
	pp, oo = processPermutation([]int{9, 1, 8, 2, 7, 3, 6, 4, 5})
	fmt.Printf("%v, %v\n", pp, oo)
	*/
	maxNum, seed := 0, 0
	var ok bool
	p := shared.NewPermutator(perm)
	perm, _ = p.NextPerm()
	for perm != nil {
		//fmt.Printf("%v\n", perm)
		if n, ok := processPermutation(perm); ok {
			p := getNum(perm)
			fmt.Printf("Found: %v, Seed: %v\n", p, n)
			if maxNum < p {
				maxNum = p
				seed = n
			}
		}
		perm, ok = p.NextPerm()
		if !ok {
			break
		}
	}
	fmt.Printf("Max Found: %v, Seed: %v\n", maxNum, seed)

}
