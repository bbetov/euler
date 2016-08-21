package main

import (
	"fmt"

	"github.com/bbetov/euler/shared/primer"
)

func isPandigital(n uint64) bool {
	var maxBit, fl uint64
	for n > 0 {
		r := n % 10
		bit := uint64(1) << r
		if r == 0 || fl&bit > 0 { // already found or zero
			return false
		}
		if maxBit < r {
			maxBit = r
		}
		fl |= bit
		n /= 10
	}

	for i := uint64(1); i <= maxBit; i++ {
		bit := uint64(1) << i
		if fl&bit == 0 {
			return false
		}
	}

	return true
}

func main() {
	// fairly slow - I need to think about how to optimize the process
	pr := primer.NewPrimerWithSegSize(uint32(10000000))
	p, _ := pr.NextPrime()
	var maxPan uint64
	var err error
	var cntr int
	for p <= 987654321 {
		cntr++
		if isPandigital(p) {
			maxPan = p
			fmt.Printf("Potential Candidate: %v\n", maxPan)
		}
		if p, err = pr.NextPrime(); err != nil {
			break
		}

		if cntr%500000 == 0 {
			println(p)
		}
	}
	fmt.Printf("Answer: %v", maxPan)
}
