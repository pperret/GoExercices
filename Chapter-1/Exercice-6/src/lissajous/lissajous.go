package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0xff, 0, 0, 0xff}, color.RGBA{0, 0xff, 0, 0xff}, color.RGBA{0, 0, 0xff, 0xff}}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles      = 5     // number of complete x oscillator revolutions
		res         = 0.001 // angular resolution
		size        = 300   // image canvas covers [size..+size]
		nframes     = 64    // number of animation frames
		delay       = 8     // delay between frames in 10ms units
		phaseOffset = 1.0   // Phase offet between colors
	)
	colors := len(palette) - 1   // number of colors
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			for c := 0; c < colors; c++ {
				x := math.Sin(t)
				y := math.Sin(t*freq + phase + phaseOffset*float64(c))
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c+1))
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
