package main

import "math/big"

func expansionGenerator() func() (*big.Int, *big.Int) {
	two := big.NewInt(2)
	n, d := big.NewInt(1), big.NewInt(2)
	return func() (*big.Int, *big.Int) {
		tmp := big.Int{}
		tmp.Set(d)
		d.Mul(d, two)
		d.Add(d, n)
		n.Set(&tmp)
		return n, d
	}
}

func getSquareRootConvergents(limit int) int {
	gen := expansionGenerator()
	n, d := big.NewInt(1), big.NewInt(2)
	var counter int
	num, den := big.NewInt(0), big.NewInt(0)
	for i := 0; i < limit; i++ {
		num.Set(d)
		num.Add(num, n)
		den.Set(d)
		if len(num.String()) > len(den.String()) {
			counter++
		}
		n, d = gen()
	}
	return counter
}

func main() {
	println(getSquareRootConvergents(1000))
}
