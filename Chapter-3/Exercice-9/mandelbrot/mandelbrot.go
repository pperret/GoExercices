// Mandembrot server
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler generating the graphic
func handler(w http.ResponseWriter, r *http.Request) {
	const (
		width, height = 1024, 1024
	)

	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}

	// Get X position
	xpos := 0.0 // Default value
	xt := r.Form["x"]
	if xt != nil {
		xx, err := strconv.ParseFloat(xt[0], 64)
		if err == nil {
			xpos = xx
		}
	}

	// Get Y position
	ypos := 0.0 // Default value
	yt := r.Form["y"]
	if yt != nil {
		yy, err := strconv.ParseFloat(yt[0], 64)
		if err == nil {
			ypos = yy
		}
	}

	// Get zoom
	zoom := 100.0 // Default value
	zt := r.Form["zoom"]
	if zt != nil {
		zz, err := strconv.ParseFloat(zt[0], 64)
		if err == nil {
			zoom = zz
		}
	}

	xmin := float64(xpos) - (2 * 100 / zoom)
	xmax := float64(xpos) + (2 * 100 / zoom)
	ymin := float64(ypos) - (2 * 100 / zoom)
	ymax := float64(ypos) + (2 * 100 / zoom)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img) // NOTE: ignoring errors
}

// Compute the color of a pixel at a specific location
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return getColor(n)
		}
	}
	return color.Black
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
