// Test package to benchmark several implementation methods
package popcount

import (
	"testing"
)

// Bench the "efficient" method
func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount1(255)
	}
}

// Bench the "unefficient" method
func BenchmarkPopcount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(255)
	}
}
