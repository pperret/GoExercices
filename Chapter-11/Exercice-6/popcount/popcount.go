// Package popcount implements several versions of bits counting in an integer
package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount1 returns the population count (number of set bits) of x.
func PopCount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCount2 is a second version using shift
func PopCount2(x uint64) int {
	n := 0
	for i := 0; i < 64; i++ {
		if (x & 1) != 0 {
			n++
		}
		x >>= 1
	}
	return n
}

// PopCount3 if an "optimized" version of bits counting
func PopCount3(x uint64) int {
	var n int
	for n = 0; x != 0; n++ {
		x &= (x - 1)
	}
	return n
}

func main() {
	fmt.Printf("PopCount1 -> %d\n", PopCount1(0))
	fmt.Printf("PopCount1 -> %d\n", PopCount2(0))
	fmt.Printf("PopCount1 -> %d\n", PopCount3(0))
}
