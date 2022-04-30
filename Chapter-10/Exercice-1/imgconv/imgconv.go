// The imgconv command reads a image (in PNG, GIF or JPEG format) from the standard input
// and writes it in another format (PNG, GIF or JPEG) to the standard output.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

// main is the entry point of the program
func main() {
	format := flag.String("fmt", "jpeg", "Output image format (png, gif or jpeg)")
	flag.Parse()

	if err := imgconv(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}

}

// imgconv converts the image into another format (or the same one)
func imgconv(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch strings.ToLower(format) {
	case "jpg", "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{NumColors: 256})
	default:
		return fmt.Errorf("Unknown output format")
	}
}
