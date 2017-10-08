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
	
	xmin := float64(xpos) - (2*100/zoom)
	xmax := float64(xpos) + (2*100/zoom)
	ymin := float64(ypos) - (2*100/zoom)
	ymax := float64(ypos) + (2*100/zoom)
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax-xmin) + xmin
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
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			var r uint8 = contrast*n
			var b uint8 = 255-contrast*n
			var g uint8 = 255-contrast* (n-10)*(n-10)
			return color.RGBA{r, g, b, 255}
		}
	}
	return color.Black
}
