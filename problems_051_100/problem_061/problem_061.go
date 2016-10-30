package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
	"github.com/bbetov/euler/shared/polynum"
)

type numStore struct {
	num, fistDig, lastDig int
}

var numbers = make([][]numStore, 6)

func init() {
	var n uint64
	for i := 3; i < 9; i++ {
		for idx := 1; ; idx++ {
			switch i {
			case 3:
				n = polynum.Triangular(idx)
			case 4:
				n = polynum.Square(idx)
			case 5:
				n = polynum.Pentagonal(idx)
			case 6:
				n = polynum.Hexagonal(idx)
			case 7:
				n = polynum.Heptagonal(idx)
			case 8:
				n = polynum.Octagonal(idx)
			}
			if n < 1000 {
				continue
			}
			if n >= 10000 {
				break
			}
			in := int(n)
			fd := in / 100
			ld := in % 100
			if fd > 0 && ld > 0 {
				numbers[i-3] = append(numbers[i-3], numStore{num: in, fistDig: fd, lastDig: ld})
			}
		}
	}
}

func checkForMatch(numbers [][]numStore) bool {
	for _, v3 := range numbers[0] {
		for _, v4 := range numbers[1] {
			if v3.lastDig != v4.fistDig {
				continue
			}
			for _, v5 := range numbers[2] {
				if v4.lastDig != v5.fistDig {
					continue
				}
				for _, v6 := range numbers[3] {
					if v5.lastDig != v6.fistDig {
						continue
					}
					for _, v7 := range numbers[4] {
						if v6.lastDig != v7.fistDig {
							continue
						}
						for _, v8 := range numbers[5] {
							if v8.lastDig == v3.fistDig && v8.fistDig == v7.lastDig {
								fmt.Printf("%v\n", []int{v3.num, v4.num, v5.num, v6.num, v7.num, v8.num})
								fmt.Printf("The sum is %v\n", v3.num+v4.num+v5.num+v6.num+v7.num+v8.num)
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
}

func main() {
	p := shared.NewPermutator([]int{0, 1, 2, 3, 4, 5})
	for {
		if n, ok := p.NextPerm(); ok {
			var combo [][]numStore
			for i := 0; i < len(n); i++ {
				combo = append(combo, numbers[n[i]])
			}
			if checkForMatch(combo) {
				break
			}
		}
	}
}
