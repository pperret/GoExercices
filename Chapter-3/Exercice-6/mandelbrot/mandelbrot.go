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
		oversampling           = 2
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
	c := color.RGBA{0, 0, 0, 255}
	n := mandelbrot(z)
	if n < 255 {
		c = getColor(n)
	}
	return c
}

// Computes the RGB color from the iteration count
func getColor(n int) color.RGBA {
	var sr, sg, sb uint8
	const contrast = 15
	n *= contrast
	switch {
	case n < 64:
		sr, sg, sb = 0, uint8(n*4), 255
	case n < 128:
		sr, sg, sb = 0, 255, uint8(255-(n-64)*4)
	case n < 192:
		sr, sg, sb = uint8((n-128)*4), 255, 0
	case n < 256:
		sr, sg, sb = 255, uint8(256-(n-191)*4), 0
	}
	return color.RGBA{sr, sg, sb, 255}
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
