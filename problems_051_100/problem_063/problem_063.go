package main

import "math/big"

func main() {
	var counter int
	for n := int64(1); n < 10; n++ {
		nn := big.NewInt(n)
		np := big.NewInt(n)
		for p := 1; ; p++ {
			if len(np.String()) < p {
				break
			}
			counter++
			np.Mul(np, nn)
		}
	}
	println(counter)
}
