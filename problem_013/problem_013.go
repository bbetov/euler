package main

// Could have used big/Int, but I thought that the point is to implement the actual numeric addition.

import (
	"bufio"
	"os"
	"strconv"
)

type Integer struct {
	digits []int
}

func (ii *Integer) add(a *Integer) {
	var b, s []int
	if len(ii.digits) > len(a.digits) {
		b = ii.digits
		s = a.digits
	} else {
		b = a.digits
		s = ii.digits
	}
	l := len(b) + 1
	tmp := make([]int, l)

	carry := 0
	for idx, val := range b {
		sm := val + carry
		if idx < len(s) {
			sm += s[idx]
		}
		carry = 0
		if sm > 9 {
			carry = 1
			sm = sm % 10
		}
		tmp[idx] = sm
	}
	if carry > 0 {
		tmp[len(b)] = carry
	}
	ii.digits = tmp
}

func (ii *Integer) set(s string) {
	for i := len(s) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		ii.digits = append(ii.digits, n)
	}
}

func (ii *Integer) String() (s string) {
	for i := len(ii.digits) - 1; i >= 0; i-- {
		s += strconv.Itoa(ii.digits[i])
	}
	return s
}

func (ii *Integer) norm() {
	var tmp []int
	zero := len(ii.digits)
	for i := len(ii.digits) - 1; i >= 0; i-- {
		if ii.digits[i] > 0 {
			zero = i
			break
		}
	}

	for idx, val := range ii.digits {
		tmp = append(tmp, val)
		if idx == zero {
			break
		}
	}
	ii.digits = tmp
}

func loadData(path string) ([]string, error) {
	var rv []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rv = append(rv, line)
	}

	return rv, scanner.Err()
}

func main() {
	lines, _ := loadData("input.txt")

	i1 := new(Integer)
	for _, val := range lines {
		i := new(Integer)
		i.set(val)
		i1.add(i)
		i1.norm()
	}
	println(i1.String()[:10])
}
