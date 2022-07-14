// Bench the several implementations of bits counting in an integer
// On my platform (darwin [MacOS] with Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz)
// PopCount1 is faster than PopCount3 when the bit count is greater than or equal to 4
package main

import (
	"flag"
	"testing"
)

var value = flag.Uint64("value", 255, "Value to process")
var expected = flag.Int("expected", 8, "Expected result")

// Bench the initial implementation
func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// The result is checked to prevent the compiler eliminating the function call.
		r := PopCount1(*value)
		if r != *expected {
			b.Fatalf("PopCount1(%d)->%d instead of %d", *value, r, *expected)
		}
	}
}

// Bench the "optimized" implémentation
func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// The result is checked to prevent the compiler eliminating the function call.
		r := PopCount2(*value)
		if r != *expected {
			b.Fatalf("PopCount2(%d)->%d instead of %d", *value, r, *expected)
		}
	}
}

// Bench the "Divide and Conquer Strategy" (really optimized) implémentation
func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// The result is checked to prevent the compiler eliminating the function call.
		r := PopCount3(*value)
		if r != *expected {
			b.Fatalf("PopCount3(%d)->%d instead of %d", *value, r, *expected)
		}
	}
}
