// Bench the several implementations of bits counting in an integer
package popcount_test

import (
	"popcount"
	"testing"
)

// Bench the initial implementation
func BenchmarkPopCount1(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		popcount.PopCount1(255)
	}
}

// Bench the "optimized" implémentation
func BenchmarkPopCount2(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		popcount.PopCount2(255)
	}
}

