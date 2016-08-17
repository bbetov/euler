package main

import "github.com/bbetov/euler/shared"

type digdiv struct {
	sd, div int // Starting digit and divisor prime
}

func getNum(s []int) (n int) {
	mul := 1
	for i := len(s) - 1; i >= 0; i-- {
		n += s[i] * mul
		mul *= 10
	}
	return n
}

func main() {
	divisors := []digdiv{
		{7, 17},
		{6, 13},
		{5, 11},
		{4, 7},
		{3, 5},
		{2, 3},
		{1, 2},
	}
	var sum uint64
	p := shared.NewPermutator([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	for {
		per, ok := p.NextPerm()
		if !ok {
			break
		}
		found := true
		for _, d := range divisors {
			n := getNum(per[d.sd : d.sd+3])
			if n%d.div != 0 {
				found = false
				break
			}
		}
		if found {
			sum += uint64(getNum(per))
		}
	}
	println(sum)
}
