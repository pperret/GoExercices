// Benchmarck of several implementations of mandelbrot program
package mandelbrot

import (
	"testing"
)

// Benchmark for complex64 version
func BenchmarkComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Complex64Mandelbrot()
	}
}

// Benchmark for complex128 version
func BenchmarkComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Complex128Mandelbrot()
	}
}

// Benchmark for big float version
func BenchmarkBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigFloatMandelbrot()
	}
}

// Benchmark for big rat version
func BenchmarkBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigRatMandelbrot()
	}
}
