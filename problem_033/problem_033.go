package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func simplify(n, d int) (fn, fd int) {
	mn := shared.GetDivisorsFreq(uint64(n))
	md := shared.GetDivisorsFreq(uint64(d))
	for nk, nv := range mn {
		if dv, ok := md[nk]; ok {
			tmp := minInt(nv, dv)
			for i := 0; i < tmp; i++ {
				n /= int(nk)
				d /= int(nk)
			}
		}
	}
	return n, d
}

func getSimplifiedFraction(n, d int) (int, int, bool) {
	n1, n2 := n%10, n/10
	d1, d2 := d%10, d/10
	if n1 == 0 || n2 == 0 || d1 == 0 || d2 == 0 {
		return 0, 0, false
	}
	var fn, fd int
	if d1 == n1 {
		fn = n2
		fd = d2
	} else if d1 == n2 {
		fn = n1
		fd = d2
	} else if d2 == n1 {
		fn = n2
		fd = d1
	} else if d2 == n2 {
		fn = n1
		fd = d1
	}
	if fn == 0 || fd == 0 {
		return 0, 0, false
	}

	fn, fd = simplify(fn, fd)
	fn2, fd2 := simplify(n, d)
	if fn == fn2 && fd == fd2 {
		return fn, fd, true
	}

	return 0, 0, false
}

func main() {
	// Find fractions
	// Simplify
	// Find Denominator
	tn, td := 1, 1
	for n := 10; n < 100; n++ {
		for d := n + 1; d < 100; d++ {
			if n2, d2, ok := getSimplifiedFraction(n, d); ok {
				fmt.Printf("Frac: %v / %v  ==>  %v / %v\n", n, d, n2, d2)
				tn *= n2
				td *= d2
			}
		}
	}
	tn, td = simplify(tn, td)
	println(td)
}
