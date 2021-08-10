// Test program for temperature conversions
package main

import (
	"fmt"

	"dummy.com/C2E1/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC) // "Brrrr! 273.15°C"

	fmt.Printf("0°C -> %v\n", tempconv.CToK(0))
	fmt.Printf("0°K -> %v\n", tempconv.KToC(0))
	fmt.Printf("0°F -> %v\n", tempconv.FToK(0))
	fmt.Printf("0°K -> %v\n", tempconv.KToF(0))

}
