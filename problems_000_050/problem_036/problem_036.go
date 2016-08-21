package main

// this converts it backwards, but since we are looking for a palindrome
// it doesn't matter
func toDigitsArr(n, base int) []int {
	rv := []int{}
	for n > 0 {
		d := n % base
		rv = append(rv, d)
		n /= base
	}
	return rv
}

func isPalindrome(d []int) bool {
	l := len(d)
	if l == 0 {
		return false
	}
	if l == 1 {
		return true
	}

	for i := 0; i < l/2; i++ {
		if d[i] != d[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	totalSum := 0
	for i := 1; i < 1000000; i++ {
		d := toDigitsArr(i, 10)
		if isPalindrome(d) {
			d2 := toDigitsArr(i, 2)
			if isPalindrome(d2) {
				totalSum += i
			}
		}
	}
	println(totalSum)
}
