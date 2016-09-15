package main

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPairs
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

const (
	Ten   = 10
	Jack  = 11
	Queen = 12
	King  = 13
	Ace   = 14
)

type card struct {
	v int // T = 10, J = 11, Q = 12, K = 13, A = 14
	s int //
}

type cards []card

type hand struct {
	c cards
	h histogram
}

type histVal struct {
	v int // Card Value
	c int // Count
}
type histogram []histVal

func (c cards) Len() int {
	return len(c)
}

func (h histogram) Len() int {
	return len(h)
}

func (c cards) Less(i, j int) bool {
	if c[i].v < c[j].v {
		return true
	}
	return false
}

func (h histogram) Less(i, j int) bool {
	if h[i].c > h[j].c {
		return true
	}

	if h[i].c == h[j].c && h[i].v > h[j].v {
		return true
	}

	return false
}

func (c cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (h histogram) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hand) isSameSuit() bool {
	suit := h.c[0].s
	for i := 1; i < len(h.c); i++ {
		if suit != h.c[i].s {
			return false
		}
	}
	return true
}

func (h hand) isConsecutive() bool {
	st := h.c[0].v
	for i := 1; i < len(h.c); i++ {
		st++
		if st != h.c[i].v {
			return false
		}
	}
	return true
}

func (h hand) getValCounts() []int {
	sl := make([]int, Ace+1)
	for _, c := range h.c {
		sl[c.v]++
	}
	return sl
}

func (h hand) value() int {
	sameSuit := h.isSameSuit()
	consecutive := h.isConsecutive()
	if consecutive && sameSuit {
		if h.c[4].v == Ace {
			return royalFlush
		}
		return straightFlush
	}
	if sameSuit {
		return flush
	}
	if consecutive {
		return straight
	}
	if h.h[0].c == 4 {
		return fourOfAKind
	}

	if h.h[0].c == 3 {
		if h.h[1].c == 2 {
			return fullHouse
		}
		return threeOfAKind
	}
	if h.h[0].c == 2 {
		if h.h[1].c == 2 {
			return twoPairs
		}
		return onePair
	}

	return highCard
}

func newCard(s string) (*card, error) {
	var c card
	if num, err := strconv.Atoi(s[:len(s)-1]); err == nil {
		c.v = num
	} else {
		switch s[0] {
		case 'T':
			c.v = Ten
		case 'J':
			c.v = Jack
		case 'Q':
			c.v = Queen
		case 'K':
			c.v = King
		case 'A':
			c.v = Ace
		default:
			return nil, errors.New("Invalid value")
		}
	}
	switch s[len(s)-1] {
	case 'C':
		c.s = 1
	case 'D':
		c.s = 2
	case 'H':
		c.s = 3
	case 'S':
		c.s = 4
	default:
		return nil, errors.New("Invalid suit")
	}
	return &c, nil
}

func newHand(inp []string) (*hand, error) {
	var h hand
	if len(inp) != 5 {
		return nil, errors.New("Invalid number of cards.")
	}

	for _, s := range inp {
		c, err := newCard(s)
		if err != nil {
			return nil, err
		}
		h.c = append(h.c, *c)
	}
	sort.Sort(h.c)
	sl := h.getValCounts()
	for i, v := range sl {
		if v > 0 {
			h.h = append(h.h, histVal{v: i, c: v})
		}
	}
	sort.Sort(h.h)
	return &h, nil
}

func processFile() (n int, err error) {
	file, err := os.Open("p054_poker.txt")
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cards := strings.Split(scanner.Text(), " ")
		h1, err1 := newHand(cards[:5])
		if err1 != nil {
			return 0, err1
		}
		h2, err2 := newHand(cards[5:])
		if err2 != nil {
			return 0, err2
		}

		v1 := h1.value()
		v2 := h2.value()
		if v1 > v2 {
			n++
			continue
		}
		if v1 == v2 { // same val
			for i := 0; i < len(h1.h); i++ {
				if h1.h[i].v > h2.h[i].v {
					n++
					break
				}
				if h1.h[i].v < h2.h[i].v {
					break
				}
			}
		}
	}
	return n, scanner.Err()
}

func main() {
	n, _ := processFile()
	println(n)
}
