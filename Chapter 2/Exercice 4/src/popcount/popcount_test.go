// Bench the several implements of bits counting
package popcount

import (
	"testing"
)

// Bench the "efficient" implementation
func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(255)
	}
}

// Bench the "unefficient" implementation
func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(255)
	}
}
