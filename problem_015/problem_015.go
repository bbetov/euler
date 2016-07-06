package main

func main() {
	// Lattice paths are binomial coefficients 2 * n choose n
	// see: https://en.wikipedia.org/wiki/Lattice_path#North-East_lattice_paths

	// Efficient binomial coefficients are calculated with the following formula
	// http://blog.plover.com/math/choose.html
	size := uint64(20)
	start := 2 * size
	paths := uint64(1)

	for i := uint64(1); i <= size; i++ {
		paths *= start
		start--
		paths /= i
	}

	println(paths)
}
