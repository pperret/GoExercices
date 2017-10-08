// Benchmarck of several implementations of echo program
package echo

import (
	"testing"
)

// Benchmark of the un-efficient version
func BenchmarkMethod1(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		Method1()
	}
}

// Benchmark of the efficient version
func BenchmarkMethod2(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		Method2()
	}
}
