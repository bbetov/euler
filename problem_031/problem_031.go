package main

var coins = []int{200, 100, 50, 20, 10, 5, 2, 1}

// https://andrew.neitsch.ca/publications/m496pres1.nb.pdf
func waysToSum(total, index int) int {
	if total < 0 || index < 0 {
		return 0
	}
	if total == 0 {
		return 1
	}
	return waysToSum(total, index-1) + waysToSum(total-coins[index], index)
}

func waysToSumBruteForce(total int) int {
	count := 0
	for a := 0; a <= total; a += 200 {
		for b := 0; b <= total; b += 100 {
			for c := 0; c <= total; c += 50 {
				for d := 0; d <= total; d += 20 {
					for e := 0; e <= total; e += 10 {
						for f := 0; f <= total; f += 5 {
							for g := 0; g <= total; g += 2 {
								if a+b+c+d+e+f+g <= total {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	return count
}

func main() {

	println(waysToSum(200, len(coins)-1))
	println(waysToSumBruteForce(200))
}
