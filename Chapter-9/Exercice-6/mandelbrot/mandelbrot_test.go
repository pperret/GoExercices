package main

import (
	"testing"
)

func BenchmarkRender1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(1, nil)
	}
}

func BenchmarkRender2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(2, nil)
	}
}

func BenchmarkRender3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(3, nil)
	}
}

func BenchmarkRender4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(4, nil)
	}
}
