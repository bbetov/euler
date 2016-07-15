package main

import "fmt"

func div(dvsr, nd int) (r []byte) {
	cur := 1
	nz, ps, of := 0, 0, 0

	for it := 0; cur > 0 && it < nd; it++ {
		if cur < dvsr {
			r = append(r, 0)
		} else {
			d := cur / dvsr
			r = append(r, byte(d))
			cur = cur % dvsr
			if nz == 0 {
				nz = len(r) - 1
				cur *= 10
				continue
			}

			if byte(d) == r[nz] && ps == 0 {
				ps = len(r) - 1
				of = 0
				if ps-nz == 1 {
					break
				}
			}

			if r[nz+of] == r[ps+of] {
				if of >= ps-nz {
					break
				}
				of++
			} else {
				ps = 0
				of = 0
			}

		}
		cur *= 10
	}
	return r[nz:ps]
}

func main() {
	digits := div(6, 200)
	fmt.Printf("%v", digits)
}
