package main

import "github.com/bbetov/euler/shared"

func getTriangularNumber(maxDivs int) (n uint64, ndiv int) {
	num := uint64(1)
	for tn := uint64(1); ; tn += num {
		divcnt := 2
		sq := shared.SqrtUInt64(tn)
		for d := uint64(2); d < sq; d++ {
			if tn%d == 0 {
				divcnt++
				if tn/d > sq {
					divcnt++
				}
			}
		}
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
