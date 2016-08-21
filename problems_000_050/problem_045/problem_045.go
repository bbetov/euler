package main

import "github.com/bbetov/euler/shared"

func getPentagonalNumber(i uint64) (n uint64) {
	return (i * (i*3 - 1)) / 2
}
func isPentagonal(P uint64) bool {
	// Compute n
	n := shared.SqrtUInt64(P*24 + 1)
	n = (n + 1) / 6
	return P == getPentagonalNumber(n)
}

func getHexagonalNumber(i uint64) (n uint64) {
	return i * (2*i - 1)
}
func isHexagonal(P uint64) bool {
	// Compute n
	n := shared.SqrtUInt64(P*8 + 1)
	n = (n + 1) / 4
	return P == getHexagonalNumber(n)
}

func getTriangleNumber(i uint64) (n uint64) {
	return (i * (i + 1)) / 2
}

func main() {
	for i := uint64(286); ; i++ {
		T := getTriangleNumber(i)
		if isHexagonal(T) && isPentagonal(T) {
			println(T)
			break
		}
	}

}
