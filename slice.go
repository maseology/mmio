// just a set of slice manipulation tools

package mmio

import (
	"math"
)

// Rev is quick function used to reverse order of a slice
func Rev(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// RevF is quick function used to reverse order of a float64 slice
func RevF(s []float64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// OnlyPositive removes all value <= 0.0 and all NaN's
func OnlyPositive(s []float64) []float64 {
	// fmt.Println(len(s))
	// fmt.Println(s)
	var x []int
	for i := range s {
		if s[i] <= 0 || math.IsNaN(s[i]) {
			x = append(x, i)
		}
	}
	Rev(x)
	for _, i := range x {
		s = append(s[:i], s[i+1:]...)
	}
	return s[:len(s)]
	// // s = append([]float64(nil), s[:len(s)-len(x)]...)
	// fmt.Println(len(s))
	// fmt.Println(cap(s))
	// fmt.Println(len(x))
	// fmt.Println(s)
}

// Sequential builds a n+1 length slice indexed from 0 to n
func Sequential(n int) []int {
	iout := make([]int, n+1)
	for i := 0; i <= n; i++ {
		iout[i] = i
	}
	return iout
}

// Distinct returns a unique slice
func Distinct(ints []int) []int {
	k := make(map[int]bool)
	for _, i := range ints {
		if _, ok := k[i]; !ok {
			k[i] = true
		}
	}
	l, ii := make([]int, len(k)), 0
	for i := range k {
		l[ii] = i
		ii++
	}
	return l
}
