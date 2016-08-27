package main

import "fmt"

func isLeap(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}

func numDaysInYear(y int) (d int) {
	if isLeap(y) {
		return 366
	}
	return 365
}

func numDaysInMonth(m, y int) (d int) {
	switch m {
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeap(y) {
			return 29
		}
		return 28
	default:
		return 31
	}
}

var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func sundaysOnFirst() int {
	// We know that Jan 1, 1900 was Monday (0)
	// Our start date is Jan 1, 1901
	d1 := numDaysInYear(1900)
	d1Off := d1 % 7
	fmt.Printf("Jan 1, 1901 ==> %s", days[d1Off])
	cnt := 0
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			if d1Off == 6 {
				cnt++
			}
			d1Off = (numDaysInMonth(m, y) + d1Off) % 7
			//fmt.Printf("%d 1, %d ==> %s", m+1, y, days[d1Off])
		}
	}
	return cnt
}

func main() {
	println(sundaysOnFirst())
}
