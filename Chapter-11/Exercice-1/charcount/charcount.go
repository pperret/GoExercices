// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Charcount computes counts of Unicode characters from a reader
func Charcount(r io.Reader) (map[rune]int, [utf8.UTFMax + 1]int, int, int, int, int, error) {
	counts := make(map[rune]int)    // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters
	letter := 0
	digit := 0
	space := 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, utflen, 0, 0, 0, 0, err
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			letter++
		}
		if unicode.IsDigit(r) {
			digit++
		}
		if unicode.IsSpace(r) {
			space++
		}
		counts[r]++
		utflen[n]++
	}
	return counts, utflen, letter, digit, space, invalid, nil
}

// main is the entry point of the program
func main() {

	counts, utflen, letter, digit, space, invalid, err := Charcount(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
	if letter > 0 {
		fmt.Printf("%d letters\n", letter)
	}
	if digit > 0 {
		fmt.Printf("%d digits\n", digit)
	}
	if space > 0 {
		fmt.Printf("%d spaces\n", space)
	}
}
