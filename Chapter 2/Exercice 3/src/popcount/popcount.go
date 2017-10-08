// Package popcount implements several methods to get the number of bits set in an integer
package popcount

// pc[i] is the population count of i.
var pc [256]byte

// Initialize pc array (number of bits set for each 8-bits value)
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
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

// Second method using a loop
func PopCount2(x uint64) int {
	n := 0
	for i:=0 ; i<64 ; i+=8 {
		n += int(pc[byte(x>>uint(i))])
	}
	return n
}