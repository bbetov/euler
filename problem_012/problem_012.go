package main

import "github.com/bbetov/euler/shared"

func main() {
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
		if divcnt > 500 {
			println(tn)
			//println(divcnt)
			break
		}
		num++
	}
}
