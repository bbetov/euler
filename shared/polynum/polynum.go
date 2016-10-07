package polynum

func Triangular(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return (n * (n + 1)) / 2
}

func Square(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return n * n
}

func Pentagonal(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return (n * (3*n - 1)) / 2
}

func Hexagonal(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return n * (2*n - 1)
}

func Heptagonal(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return (n * (5*n - 1)) / 2
}

func Oxagonal(index int) uint64 {
	if index <= 0 {
		return 0
	}
	n := uint64(index)
	return n * (3*n - 2)
}
