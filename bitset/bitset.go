package bitset

type BitsetNumber interface {
	~uint | ~uint32 | ~uint64 | ~uint16 | ~uint8
}

// T is the container type, B is the type to identify the bits.
// First bit is bit 0.
type BitSet[T BitsetNumber, B BitsetNumber] struct {
	value T
}

func BitsetValue[T BitsetNumber, B BitsetNumber](bitsToSet ...B) BitSet[T, B] {
	bitset := BitSet[T, B]{}
	for _, b := range bitsToSet {
		bitset.Set(b)
	}
	return bitset
}

// First bit is bit 0.
func (bitset *BitSet[T, B]) Set(bit B) {
	bitset.value = bitset.value | (1 << bit)
}

// First bit is bit 0.
func (bitset *BitSet[T, B]) Clear(bit B) {
	bitset.value = bitset.value & ^(1 << bit)
}

func (bitset *BitSet[T, B]) SetAll() {
	bitset.value = bitset.value | ^(T(0))
}

func (bitset *BitSet[T, B]) ClearAll() {
	bitset.value = 0
}

func (bitset *BitSet[T, B]) Test(bit B) bool {
	return ((bitset.value & (1 << bit)) >> bit) == 1
}

func (bitset *BitSet[T, B]) TestAll(bits ...B) bool {
	tester := T(0)
	for bit := range bits {
		tester = tester | (1 << T(bit))
	}
	return (bitset.value & tester) > 0
}

func (bitset *BitSet[T, B]) Equal(b2 BitSet[T, B]) bool {
	return bitset.value == b2.value
}

func (bitset *BitSet[T, B]) Union(b2 BitSet[T, B]) {
	bitset.value = bitset.value | b2.value
}

// Sets bitset to only the values that both share.
func (bitset *BitSet[T, B]) InnerJoin(b2 BitSet[T, B]) {
	bitset.value = bitset.value & b2.value
}
