package main

import (
	"fmt"
)

// Number of bytes per unit
const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("KB=%T, %[1]v\n", KB)

	i := YB / 1000000000
	fmt.Printf("YB=%T, %[1]v\n", i)
}
