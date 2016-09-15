package main

import (
	"strings"
	"testing"
)

var tests = []struct {
	i           string
	hl, c0, val int
}{
	{"5H KS 5C 7D 9H", 4, King, onePair},
	{"AH KS 7C 7D 7H", 3, Ace, threeOfAKind},
}

func TestCards(t *testing.T) {
	for _, tt := range tests {
		cds := strings.Split(tt.i, " ")
		h, _ := newHand(cds)
		if len(h.h) != tt.hl {
			t.Error("Expected hist len ", tt.hl, " got ", len(h.h))
		}
		if h.c[4].v != tt.c0 {
			t.Error("Expected highest card ", tt.c0, " got ", h.c[0].v)
		}
		if h.value() != tt.val {
			t.Error("Expected highest card ", tt.val, " got ", h.value())
		}
	}
}
