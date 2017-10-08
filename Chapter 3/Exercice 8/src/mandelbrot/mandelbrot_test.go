// Benchmarck of several implementations of mandelbrot program
package mandelbrot

import (
	"testing"
)

// Benchmark for complex64 version
func BenchmarkComplex64(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		MandelbrotComplex64()
	}
}


// Benchmark for complex128 version
func BenchmarkComplex128(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		MandelbrotComplex128()
	}
}

// Benchmark for big float version
func BenchmarkBigFloat(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		MandelbrotBigFloat()
	}
}

// Benchmark for big rat version
func BenchmarkBigRat(b *testing.B) {
	for i:=0 ; i<b.N ; i++ {
		MandelbrotBigRat()
	}
}
