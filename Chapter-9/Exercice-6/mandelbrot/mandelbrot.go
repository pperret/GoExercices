// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"sync"
)

// workerRequest contains data of a request to perform by workers
// (the computation of one pixel of the image)
type workerRequest struct {
	x int
	y int
	z complex128
}

// main is the entry point of the program
func main() {
	// 3 is the most efficient value for me (core i9)
	render(3, os.Stdout)
}

// render computes the mandelbrot image
func render(workerNumber int, writer io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	// Create a channel to send requests to workers
	requestChannel := make(chan workerRequest, width*height)

	// Create the image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Create a WaitGroup to track go runtines completion
	var wg sync.WaitGroup

	// Start the workers
	// Workers stop when the channel is empty and closed
	for i := 0; i < workerNumber; i++ {
		wg.Add(1)
		go func(in <-chan workerRequest) {
			for req := range in {
				img.Set(req.x, req.y, mandelbrot(req.z))
			}
			wg.Done()
		}(requestChannel)
	}

	// Send requests to workers for each point
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			requestChannel <- workerRequest{px, py, z}
		}
	}

	// Close the request channel.
	// Workers will stop when all the points will be computed.
	close(requestChannel)

	// Wait for workers completion
	wg.Wait()

	// Encode the image (if a writer is provided)
	if writer != nil {
		png.Encode(writer, img) // NOTE: ignoring errors
	}
}

// Compute the color of a pixel at a specific location
func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
