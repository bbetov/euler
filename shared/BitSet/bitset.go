package BitSet

const (
	allBits uint64 = uint64(0xFFFFFFFFFFFFFFFF)
)

// BitSet represents a simple bitset
type bitSet struct {
	storage []uint64
}

// BitSet is simply an array of booleans packed as uint64 numbers (i.e. 64 bits per elements)
type BitSet interface {
	// Set sets/unsets a specific bit at a position which. Resizes if necessary.
	Set(which uint32, val bool)
	// Reset resets all bits in the array to either 1 or 0 based on val
	Reset(val bool)
	// IsSet checks if the bit at position which is set
	IsSet(which uint32) bool
}

// Set sets a bit to 1
func (b *bitSet) Set(which uint32, val bool) {
	bn := which / 64
	for uint32(len(b.storage)) <= bn {
		b.storage = append(b.storage, 0)
	}

	whichBit := which % 64
	if val {
		b.storage[bn] = b.storage[bn] | (uint64(1) << whichBit)
	} else {
		b.storage[bn] = b.storage[bn] &^ (uint64(1) << whichBit)
	}
}

func (b *bitSet) IsSet(which uint32) bool {
	bn := which / 64
	if uint32(len(b.storage)) <= bn {
		return false
	}

	whichBit := which % 64
	return (b.storage[bn] & (uint64(1) << whichBit)) > 0
}

// New creates a BitSet with a specific size (numbits) and sets them to a value
func New(numbits uint32, value bool) BitSet {
	bn := numbits / 64
	bs := &bitSet{storage: make([]uint64, bn+1)}

	if value {
		for i := uint32(0); i <= bn; i++ {
			bs.storage[i] = allBits
		}
	}
	return bs
}

func (b *bitSet) Reset(val bool) {
	fv := allBits
	if !val {
		fv = 0
	}
	for i := range b.storage {
		b.storage[i] = fv
	}
}
