package primer

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestPrimes10K(t *testing.T) {
	// Load test data
	file, err := os.Open("10kPrimes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewScanner(file)
	p := make(map[uint64]bool)
	for r.Scan() {
		line := r.Text()
		i, _ := strconv.Atoi(line)
		p[uint64(i)] = true
	}

	pr := NewPrimer()
	cntr := 0
	for cntr < len(p) {
		np, err := pr.NextPrime()
		if err != nil {
			t.Errorf("Error %v.", err)
		}
		if _, ok := p[np]; !ok {
			t.Errorf("Unexpected number %v.", np)
		}
		cntr++
	}
}
