package mmio

import "math"

// BitArray converts a slice of bytes to a slice of boolean
func BitArray(b []byte) []bool {
	a := make([]bool, len(b)*8)
	for i, v := range b {
		ba := BitArray1(v)
		for p := 0; p < 8; p++ {
			a[i*8+p] = ba[p]
		}
	}
	return a
}

// BitArray1 converts a byte into a slice of boolean
func BitArray1(b byte) []bool {
	a := make([]bool, 8)
	if b&1 == 1 {
		a[0] = true
	}
	if b&2 == 2 {
		a[1] = true
	}
	if b&4 == 4 {
		a[2] = true
	}
	if b&8 == 8 {
		a[3] = true
	}
	if b&16 == 16 {
		a[4] = true
	}
	if b&32 == 32 {
		a[5] = true
	}
	if b&64 == 64 {
		a[6] = true
	}
	if b&128 == 128 {
		a[7] = true
	}
	return a
}

// BitArrayRev converts a slice of boolean to a slice of bytes
func BitArrayRev(b []bool) []byte {
	n := int(math.Ceil(float64(len(b)) / 8))
	a := make([]byte, 0, n)
	for i := 0; i < len(b); i += 8 {
		i8 := i + 8
		if i8 > len(b) {
			i8 = len(b)
		}
		bb := b[i:i8]
		for i2, j := 0, len(bb)-1; i2 < j; i2, j = i2+1, j-1 { // reversing slice
			bb[i2], bb[j] = bb[j], bb[i2]
		}
		a = append(a, BitArray1Rev(bb))
	}
	return a
}

// BitArray1Rev converts a slice of boolean to a byte
// https://stackoverflow.com/questions/73710132/golang-convert-8bool-to-byte
func BitArray1Rev(b []bool) byte {
	var result byte
	for _, b := range b {
		result <<= 1
		if b {
			result |= 1
		}
	}
	return result
}
