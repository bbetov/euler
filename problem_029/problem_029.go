package main

import (
	"fmt"

	"github.com/bbetov/euler/shared"
)

var fl = [102][102]map[uint64]int{}
var pdc = make([]map[uint64]int, 102)

// key only match
func keyMatch(a map[uint64]int, b map[uint64]int) bool {
	if len(a) != len(b) || a == nil || b == nil {
		return false
	}

	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

func valueMatch(a map[uint64]int, b map[uint64]int) bool {
	for k, v := range a {
		if v2, ok := b[k]; ok {
			if v != v2 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func hasMatch(mp map[uint64]int, current int) bool {
	for i := 2; i < current; i++ {
		if keyMatch(mp, pdc[i]) {
			for po := 2; po <= 100; po++ {
				if valueMatch(mp, fl[i][po]) {
					return true
				}
			}
		}
	}
	return false
}

func makeMap(base map[uint64]int, po int) (nm map[uint64]int) {
	nm = make(map[uint64]int)
	for k, v := range base {
		nm[k] = v * po
	}
	return nm
}

func process(i int) int {
	var counter int
	for po := 2; po <= 100; po++ {
		nm := makeMap(pdc[i], po)

		if !hasMatch(nm, i) {
			fl[i][po] = nm
			counter++
		}
	}

	return counter
}

func main() {
	// 1. Get all prime decompositions 2 - 100
	for i := 2; i <= 100; i++ {
		dv := shared.GetDivisorsFreq(uint64(i))
		pdc[i] = dv
		//fmt.Printf("%v ==> %v\n", i, dv)
	}
	var counter int
	for i := 2; i <= 100; i++ {
		counter += process(i)
	}
	fmt.Printf("Total distinct numbers: %v", counter)
}
