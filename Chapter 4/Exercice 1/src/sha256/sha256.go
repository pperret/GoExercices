// The sha256 program computes the number of different bits in two SHA256 hashes (an array).
package main

import (
	"fmt"
	"crypto/sha256"
	)

func different(c1, c2 [32]uint8) int {
	n := 0
	for i := range c1 {
		for b := c1[i] ^ c2[i] ; b != 0 ; b &= (b - 1) {
			n++
		}
	}
	return n
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	d := different(c1, c2)
	fmt.Printf("%x\n%x\n%d\n", c1, c2, d)	
}