package main

import "strings"

var sub10 = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
var sub20 = []string{"eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen",
	"eighteen", "nineteen"}

var tens = []string{"ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

func getSub100(num int) (s string) {
	if num <= 10 {
		return sub10[num]
	}
	if num < 20 {
		s = sub20[num-10-1]
		return s
	}
	if num < 100 {
		s = tens[num/10-1]
		if num%10 > 0 {
			s += " " + sub10[num%10]
		}
		return s
	}
	return s
}

func numberInWords(num int) (s string) {
	if num < 100 {
		return getSub100(num)
	}
	if num < 1000 {
		s = sub10[num/100] + " hundred"
		rem := num % 100
		if rem > 0 {
			s += " and " + getSub100(rem)
		}
	}
	if num == 1000 {
		s = "one thousand"
	}
	return s
}

func onetoThousandNumLetters() int {
	var sum int
	for i := 1; i <= 1000; i++ {
		s := numberInWords(i)
		s = strings.Replace(s, " ", "", -1)

		sum += len(s)
	}
	return sum
}

func main() {
	println(onetoThousandNumLetters())
}
