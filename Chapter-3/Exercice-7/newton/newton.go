// Newton emits a PNG image of the Newton fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, newton(z))
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// Compute the color of a pixel at a specific location
func newton(z complex128) color.Color {
	const iterations = 20
	const contrast = 15
	const eps = 1e-6
	const r1 complex128 = complex(1, 0)
	const r2 complex128 = complex(-1, 0)
	const r3 complex128 = complex(0, 1)
	const r4 complex128 = complex(0, -1)
	//var v complex128
	var n uint8
	for n = 0; n < iterations; n++ {
		z = z - (z*z*z*z-1.0)/(4*z*z*z)
		if cmplx.Abs(z*z*z*z-1.0) < eps {
			break
		}
	}

	if n < iterations {
		if cmplx.Abs(z-r1) < eps {
			return color.RGBA{0, 0, 255 - n*contrast, 255}
		} else if cmplx.Abs(z-r2) < eps {
			return color.RGBA{0, 255 - n*contrast, 0, 255}
		} else if cmplx.Abs(z-r3) < eps {
			return color.RGBA{255 - n*contrast, 0, 0, 255}
		} else if cmplx.Abs(z-r4) < eps {
			return color.Gray{255 - n*contrast}
		}
	}
	return color.Black
}
