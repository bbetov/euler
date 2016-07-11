package shared

import "strconv"

type Integer struct {
	digits []byte
}

func (ii *Integer) SumDigits() (s uint64) {
	for _, i := range ii.digits {
		s += uint64(i)
	}
	return s
}

func (ii *Integer) Add(a *Integer) {
	var b, s []byte
	if len(ii.digits) > len(a.digits) {
		b = ii.digits
		s = a.digits
	} else {
		b = a.digits
		s = ii.digits
	}
	l := len(b) + 1
	tmp := make([]byte, l)

	carry := byte(0)
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
	ii.Norm()
}

func (ii *Integer) Set(s string) {
	ii.digits = []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(s[i]))
		ii.digits = append(ii.digits, byte(n))
	}
}

func (ii *Integer) Swap(i *Integer) {
	ii.digits, i.digits = i.digits, ii.digits
}

func (ii *Integer) NumDigits() int {
	return len(ii.digits)
}

func (ii *Integer) String() (s string) {
	for i := len(ii.digits) - 1; i >= 0; i-- {
		s += strconv.Itoa(int(ii.digits[i]))
	}
	return s
}

func (ii *Integer) Norm() {
	for i := len(ii.digits) - 1; i >= 0; i-- {
		if ii.digits[i] > 0 {
			ii.digits = ii.digits[:i+1]
			return
		}
	}
}

func (ii *Integer) Multiply(m *Integer) {
	// Will start with a very simple multiplication and then improve
	curr := ii.digits
	ii.digits = []byte{}
	tmpi := Integer{}
	for i, v := range m.digits {
		tmp := make([]byte, i)
		carry := byte(0)
		for _, p := range curr {
			t := p*v + carry
			carry = t / 10
			t = t % 10
			tmp = append(tmp, t)
		}
		if carry > 0 {
			tmp = append(tmp, carry)
		}
		tmpi.digits = tmp
		ii.Add(&tmpi)
	}
}
