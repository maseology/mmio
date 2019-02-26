package mmio

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
