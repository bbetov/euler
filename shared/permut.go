package shared

type npermutator struct {
	perm []int
	// true to the right, false to the left
	dir []bool
}

type NPermutator interface {
	NextPerm() (perm []int, ok bool)
}

func (n npermutator) isMobile(index int) bool {
	if n.dir[index] {
		if index == len(n.dir)-1 || n.perm[index] < n.perm[index+1] {
			return false
		}
	} else {
		if index == 0 || n.perm[index] < n.perm[index-1] {
			return false
		}
	}
	return true
}

func (n npermutator) swap(i1, i2 int) {
	n.perm[i1], n.perm[i2] = n.perm[i2], n.perm[i1]
	n.dir[i1], n.dir[i2] = n.dir[i2], n.dir[i1]
}

func (n npermutator) NextPerm() (p []int, ok bool) {
	// Find largest mobile number
	mobilePos := -1
	var prevMobile int

	//find the largest mobile integer k
	for i := 0; i < len(n.perm); i++ {
		if n.isMobile(i) {
			if prevMobile < n.perm[i] {
				mobilePos = i
				prevMobile = n.perm[i]
			}
		}
	}
	if mobilePos < 0 {
		return nil, false
	}

	//swap k and the adjacent integer it is looking at
	if n.dir[mobilePos] {
		n.swap(mobilePos, mobilePos+1)
	} else {
		n.swap(mobilePos, mobilePos-1)
	}
	//reverse the direction of all integers larger than k
	for i := 0; i < len(n.perm); i++ {
		if n.perm[i] > prevMobile {
			n.dir[i] = !n.dir[i]
		}
	}
	return n.perm, true
}

func NewPermutator(numbers []int) NPermutator {
	p := &npermutator{}
	p.perm = numbers
	p.dir = make([]bool, len(numbers))
	return p
}
