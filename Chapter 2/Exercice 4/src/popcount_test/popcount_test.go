// Bench the several implements of bits counting
package popcount_test

import (
	"popcount"
	"testing"
)

// Bench the "efficient" implementation
func BenchmarkPopCount1(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		popcount.PopCount1(255)
	}
}

// Bench the "unefficient" implementation
func BenchmarkPopcount2(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		popcount.PopCount2(255)
	}
}
