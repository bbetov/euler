package main

import "math"

// A palindromic number reads the same both ways. The largest palindrome made
// from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func isPalindrome(num int) bool {
	if num < 0 {
		return false
	}

	number := num
	numDigits := 0

	for number != 0 {
		number /= 10
		numDigits++
	}

	if numDigits < 2 {
		return true
	}

	upperBound := numDigits / 2
	firstDigitDivisor := int(math.Pow10(numDigits - 1))

	for i := 1; i <= upperBound; i++ {
		lastDigit := num % 10
		firstDigit := num - num%firstDigitDivisor

		if lastDigit*firstDigitDivisor != firstDigit {
			return false
		}

		numDigits -= 2
		num -= firstDigit
		num /= 10
		firstDigitDivisor /= 100
	}
	return true
}

func main() {
	maxPali := 0
	for i := 1; i < 1000; i++ {
		for j := 1; j < 1000; j++ {
			n := i * j
			if isPalindrome(n) && n > maxPali {
				maxPali = n
			}
		}
	}
	println(maxPali)
}
