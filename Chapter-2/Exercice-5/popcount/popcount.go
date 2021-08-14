// Package popcount implements several versions of bits counting in an integer
package popcount

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

// PopCount2 if the "optimized" version of bits counting
func PopCount2(x uint64) int {
	var n int
	for n = 0; x != 0; n++ {
		x &= (x - 1)
	}
	return n
}

// PopCount3 if the "Divide and Conquer Strategy" version of bits counting
func PopCount3(x uint64) int {
	const k1 = 0x5555555555555555  /*  -1/3   */
	const k2 = 0x3333333333333333  /*  -1/5   */
	const k4 = 0x0f0f0f0f0f0f0f0f  /*  -1/17  */
	const kf = 0x0101010101010101  /*  -1/255 */
	x = x - ((x >> 1) & k1)        /* put count of each 2 bits into those 2 bits */
	x = (x & k2) + ((x >> 2) & k2) /* put count of each 4 bits into those 4 bits */
	x = (x + (x >> 4)) & k4        /* put count of each 8 bits into those 8 bits */
	x = (x * kf) >> 56             /* returns 8 most significant bits of x + (x<<8) + (x<<16) + (x<<24) + ...  */
	return int(x)
}
