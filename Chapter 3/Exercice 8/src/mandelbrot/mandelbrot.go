// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package mandelbrot

import (
	"math/cmplx"
	"math/big"
)

// Implementation based on complex64
func MandelbrotComplex64() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
		iterations = 200
	)
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax-xmin) + xmin
			var z complex64 = complex64(complex(x, y))
			var v complex64
			for n := 0 ; n < iterations ; n++ {
				v = v * v + z
				if cmplx.Abs(complex128(v)) > 2 {
					break 
				}
			}
		}
	}
}

// Implementation based on complex128
func MandelbrotComplex128() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
		iterations = 200
	)
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax-xmin) + xmin
			var z complex128 = complex(x, y)
			var v complex128
			for n := 0 ; n < iterations ; n++ {
				v = v * v + z
				if cmplx.Abs(v) > 2 {
					break 
				}
			}
		}
	}
}

// Implementation based on big float
func MandelbrotBigFloat() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
		iterations = 200
	)
	var c2, c4 big.Float
	c2.SetFloat64(2.0)
	c4.SetFloat64(4.0)
	var z_x, z_y big.Float
	for py := 0; py < height; py++ {
		z_y.SetFloat64(float64(py) / height * (ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			z_x.SetFloat64(float64(px) / width * (xmax-xmin) + xmin)

			var v_x, v_y big.Float
			for n := 0 ; n < iterations ; n++ {
				var v1, v2, v3, v4, v5 big.Float
				v1.Mul(&v_x, &v_x)
				v2.Mul(&v_y, &v_y)
				v3.Sub(&v1, &v2)

				v4.Mul(&v_x, &v_y)
				v5.Mul(&v4, &c2)

				v_x.Add(&v3, &z_x)
				v_y.Add(&v5, &z_y)

				// Compute modulus
				v1.Mul(&v_x, &v_x)
				v2.Mul(&v_y, &v_y)
				v3.Add(&v1, &v2)
				if v3.Cmp(&c4) > 0 {
					break 
				}
			}
		}
	}
}

// Implementation based on big rat
func MandelbrotBigRat() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height = 1024, 1024
		iterations = 200
	)
	var c2, c4 big.Rat
	c2.SetFloat64(2.0)
	c4.SetFloat64(4.0)
	var z_x, z_y big.Rat
	for py := 0; py < height; py++ {
		z_y.SetFloat64(float64(py) / height * (ymax-ymin) + ymin)
		for px := 0; px < width; px++ {
			z_x.SetFloat64(float64(px) / width * (xmax-xmin) + xmin)

			var v_x, v_y big.Rat
			for n := 0 ; n < iterations ; n++ {
				var v1, v2, v3, v4, v5 big.Rat
				v1.Mul(&v_x, &v_x)
				v2.Mul(&v_y, &v_y)
				v3.Sub(&v1, &v2)

				v4.Mul(&v_x, &v_y)
				v5.Mul(&v4, &c2)

				v_x.Add(&v3, &z_x)
				v_y.Add(&v5, &z_y)

				// Compute modulus
				v1.Mul(&v_x, &v_x)
				v2.Mul(&v_y, &v_y)
				v3.Add(&v1, &v2)
				if v3.Cmp(&c4) > 0 {
					break 
				}
			}
		}
	}
}
