package main

import "fmt"

func maxFirst(a, b, c int) (max, p, q int) {
	if a > b && a > c {
		return a, b, c
	} else if b > a && b > c {
		return b, a, c
	} else {
		return c, a, b
	}
}

func main() {
	var maxSolP, maxSol int
	for p := 10; p <= 1000; p++ {
		var sol int
		for a := 1; a < p; a++ {
			for b := 1; b < p; b++ {
				if a+b >= p {
					break
				}
				c := p - a - b
				o, p, q := maxFirst(a, b, c)
				if o*o == p*p+q*q {
					sol++
				}
			}
		}
		if maxSol < sol {
			maxSol = sol
			maxSolP = p
		}
	}
	fmt.Printf("MaxP: %v, NumSol: %v", maxSolP, maxSol)
}
