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
		contrast               = 15
	)
	offX := float64(xmax-xmin) / width / 2
	offY := float64(ymax-ymin) / height / 2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			n1 := mandelbrot(complex(x, y))
			n2 := mandelbrot(complex(x+offX, y))
			n3 := mandelbrot(complex(x, y+offY))
			n4 := mandelbrot(complex(x+offX, y+offY))
			n := (n1 + n2 + n3 + n4) / 4
			c := color.RGBA{0, 0, 0, 255}
			if n < 255 {
				r := uint8(contrast * n)
				b := uint8(255 - contrast*n)
				g := uint8(255 - contrast*(n-10)*(n-10))
				c = color.RGBA{r, g, b, 255}
			}

			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}

	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// Compute the color of a pixel at a specific location
func mandelbrot(z complex128) int {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return n
		}
	}
	return 255
}

func getcolors(img *image.RGBA, px int, py int) (uint8, uint8, uint8) {
	c := img.RGBAAt(px, py)
	return c.R, c.G, c.B
}

func oversampling2(img *image.RGBA, px1 int, py1 int, px2 int, py2 int) color.Color {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if px2 >= width || py2 >= height {
		return img.At(px1, py1)
	}

	r1, g1, b1 := getcolors(img, px1, py1)
	r2, g2, b2 := getcolors(img, px2, py2)

	r := uint8((uint(r1) + uint(r2)) / 2)
	g := uint8((uint(g1) + uint(g2)) / 2)
	b := uint8((uint(b1) + uint(b2)) / 2)

	return color.RGBA{r, g, b, 255}
}

func oversampling(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	img2 := image.NewRGBA(image.Rect(0, 0, 2*width, 2*height))

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			img2.Set(2*px, 2*py, img.At(px, py))
			img2.Set(2*px+1, 2*py, oversampling2(img, px, py, px+1, py))
			img2.Set(2*px, 2*py+1, oversampling2(img, px, py, px, py+1))
			img2.Set(2*px+1, 2*py+1, oversampling2(img, px, py, px+1, py+1))
		}
	}

	return img2
}
