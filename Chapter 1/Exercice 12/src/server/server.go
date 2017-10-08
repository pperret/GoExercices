// Web server delivering Lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler for the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	if err := r.ParseForm(); err != nil {
		log.Println(err)
	} else {
		cycles := 5.0 // Default value
		ct := r.Form["cycles"]
		if ct == nil {
			log.Println("No cycles parameter")
		} else {
			c := ct[0]
			if c == "" {
				log.Println("Cycles parameter is empty")
			} else {
				cs,errc := strconv.ParseFloat(c, 10)
				if errc != nil {
					log.Println("Invalid cycles parameter")
				} else {
					cycles = cs
					log.Println("cycles=", cycles)
				}
			}
		}
		lissajous(w, cycles)
	}
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

// Build Lissajous figure
func lissajous(out io.Writer, cycles float64) {
	const (
		res = 0.001 // angular resolution
		size = 100 // image canvas covers [size..+size]
		nframes = 64 // number of animation frames
		delay =8 // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}