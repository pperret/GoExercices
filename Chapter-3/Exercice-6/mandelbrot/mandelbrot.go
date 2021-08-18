// Mandelbrot emits a PNG image of the Mandelbrot fractal.
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
		oversampling           = 1
	)
	offX := float64(xmax-xmin) / width / oversampling
	offY := float64(ymax-ymin) / height / oversampling
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			var cl []color.Color
			for ox := 0; ox < oversampling; ox++ {
				for oy := 0; oy < oversampling; oy++ {
					cl = append(cl, colorize(complex(x+float64(ox)*offX, y+float64(oy)*offY)))
				}
			}
			c := average(cl)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// Computes the average of several colors
func average(cl []color.Color) color.Color {
	var rl, gl, bl uint64
	for _, c := range cl {
		r, g, b, _ := c.RGBA()
		rl += uint64(r)
		gl += uint64(g)
		bl += uint64(b)
	}
	n := uint64(len(cl))
	r := uint8(rl / n)
	g := uint8(gl / n)
	b := uint8(bl / n)
	return color.RGBA{r, g, b, 255}
}

// Computes the color of a pixel at a specific location
func colorize(z complex128) color.Color {
	const contrast = 15
	c := color.RGBA{0, 0, 0, 255}
	n := mandelbrot(z)
	if n < 255 {
		r := uint8(contrast * n)
		b := uint8(255 - contrast*n)
		g := uint8(255 - contrast*(n-10)*(n-10))
		c = color.RGBA{r, g, b, 255}
	}
	return c
}

// Computes the mandelbrot value of a pixel at a specific location
func mandelbrot(z complex128) int {
	const iterations = 200
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return 255
}
