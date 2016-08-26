package main

import "github.com/bbetov/euler/shared"

// Use prime factorization and determine all divisors from it
func getTriangularNumber(maxDivs int) (n uint64, ndiv int) {
	num := uint64(1)
	primes := shared.GetPrimesUInt64(100000)
	for tn := uint64(1); ; tn += num {
		divcnt := 2
		sq := shared.SqrtUInt64(tn)
		if sq > primes[len(primes)-1] {
			primes = shared.GetPrimesUInt64(sq * 10)
		}
		pd := make(map[uint64]int)
		for _, pr := range primes {
			if pr >= tn {
				break
			}
			for dv := tn; dv%pr == 0; dv /= pr {
				pd[pr]++
			}
		}
		divcnt = 1
		for _, v := range pd {
			divcnt *= v + 1
		}
		if divcnt == 1 {
			divcnt++
		}

		//fmt.Printf("%d ==> %d \n", tn, divcnt)
		if divcnt > maxDivs {
			n, ndiv = tn, divcnt
			break
		}
		num++
	}
	return n, ndiv
}

func main() {
	num, _ := getTriangularNumber(500)
	println(num)
}
