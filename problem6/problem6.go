package main

func getDiff(maxVal int) uint64 {
	mv := uint64(maxVal)
	var diff uint64
	for i := uint64(1); i <= mv; i++ {
		for j := uint64(1); j <= mv; j++ {
			if i != j {
				diff += i * j
			}
		}
	}
	return diff
}

func main() {
	println(getDiff(100))
}
