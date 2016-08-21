package main

import "math/big"

func main() {
	sum := big.NewInt(1)
	for i := int64(2); i <= 1000; i++ {
		tmp := new(big.Int).Exp(big.NewInt(i), big.NewInt(i), nil)
		sum.Add(sum, tmp)
	}
	s := sum.String()
	println(s[len(s)-10:])
}
