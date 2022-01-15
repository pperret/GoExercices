//
package popcount

import "sync"

// pc[i] is the population count of i.
var pc [256]byte

// init initializes the pc array at startup
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

var popOnce sync.Once
var pcOnce [256]byte

// initOnce initializes the pcOnce at its first use
func initOnce() {
	for i := range pcOnce {
		pcOnce[i] = pcOnce[i/2] + byte(i&1)
	}
}

// PopCountOne returns the population count (number of set bits) of x.
func PopCountOnce(x uint64) int {
	popOnce.Do(initOnce)
	return int(pcOnce[byte(x>>(0*8))] +
		pcOnce[byte(x>>(1*8))] +
		pcOnce[byte(x>>(2*8))] +
		pcOnce[byte(x>>(3*8))] +
		pcOnce[byte(x>>(4*8))] +
		pcOnce[byte(x>>(5*8))] +
		pcOnce[byte(x>>(6*8))] +
		pcOnce[byte(x>>(7*8))])
}
