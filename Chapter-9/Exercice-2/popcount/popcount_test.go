package popcount

import (
	"testing"
)

// -- Benchmarks --

// BenchmarkPopCount benchs the original version
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

// BenchmarkPopCountOnce benchs the version using sync.Once
func BenchmarkPopCountOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountOnce(0x1234567890ABCDEF)
	}
}

// GO 1.17.5
// cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
// BenchmarkPopCount-16        	1000000000	         0.2310 ns/op
// BenchmarkPopCountOnce-16    	372554997	         3.172 ns/op
