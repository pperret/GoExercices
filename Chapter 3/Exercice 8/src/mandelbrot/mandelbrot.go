// Package mandelbrot emits a PNG image of the Mandelbrot fractal.
package mandelbrot

import (
	"math/big"
	"math/cmplx"
)

// Complex64Mandelbrot is the implementation based on complex64
func Complex64Mandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
	)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex64(complex(x, y))
			var v complex64
			for n := 0; n < iterations; n++ {
				v = v*v + z
				if cmplx.Abs(complex128(v)) > 2 {
					break
				}
			}
		}
	}
}

// Complex128Mandelbrot is the implementation based on complex128
func Complex128Mandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
	)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			var v complex128
			for n := 0; n < iterations; n++ {
				v = v*v + z
				if cmplx.Abs(v) > 2 {
					break
				}
			}
		}
	}
}

// BigFloatMandelbrot is the implementation based on big float
func BigFloatMandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
	)
	var c2, c4 big.Float
	c2.SetFloat64(2.0)
	c4.SetFloat64(4.0)
	var zX, zY big.Float
	for py := 0; py < height; py++ {
		zY.SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			zX.SetFloat64(float64(px)/width*(xmax-xmin) + xmin)

			var vX, vY big.Float
			for n := 0; n < iterations; n++ {
				var v1, v2, v3, v4, v5 big.Float
				v1.Mul(&vX, &vX)
				v2.Mul(&vY, &vY)
				v3.Sub(&v1, &v2)

				v4.Mul(&vX, &vY)
				v5.Mul(&v4, &c2)

				vX.Add(&v3, &zX)
				vY.Add(&v5, &zY)

				// Compute modulus
				v1.Mul(&vX, &vX)
				v2.Mul(&vY, &vY)
				v3.Add(&v1, &v2)
				if v3.Cmp(&c4) > 0 {
					break
				}
			}
		}
	}
}

// BigRatMandelbrot is the implementation based on big rat
func BigRatMandelbrot() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
		iterations             = 200
	)
	var c2, c4 big.Rat
	c2.SetFloat64(2.0)
	c4.SetFloat64(4.0)
	var zX, zY big.Rat
	for py := 0; py < height; py++ {
		zY.SetFloat64(float64(py)/height*(ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			zX.SetFloat64(float64(px)/width*(xmax-xmin) + xmin)

			var vX, vY big.Rat
			for n := 0; n < iterations; n++ {
				var v1, v2, v3, v4, v5 big.Rat
				v1.Mul(&vX, &vX)
				v2.Mul(&vY, &vY)
				v3.Sub(&v1, &v2)

				v4.Mul(&vX, &vY)
				v5.Mul(&v4, &c2)

				vX.Add(&v3, &zX)
				vY.Add(&v5, &zY)

				// Compute modulus
				v1.Mul(&vX, &vX)
				v2.Mul(&vY, &vY)
				v3.Add(&v1, &v2)
				if v3.Cmp(&c4) > 0 {
					break
				}
			}
		}
	}
}
