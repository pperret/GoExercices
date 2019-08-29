package main

import (
	"fmt"
	"lenconv"
	"os"
	"strconv"
	"tempconv"
	"weightconv"
)

func conv(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))

	m := lenconv.Meters(t)
	fi := lenconv.Feet(t)
	fmt.Printf("%s = %s, %s = %s\n", m, lenconv.MToF(m), fi, lenconv.FToM(fi))

	k := weightconv.Kilograms(t)
	l := weightconv.Pounds(t)
	fmt.Printf("%s = %s, %s = %s\n", k, weightconv.KToL(k), l, weightconv.LToK(l))
}

func main() {
	if len(os.Args) == 1 {
		var t float64
		for {
			fmt.Printf("Enter number:")
			n, err := fmt.Scanf("%f", &t)
			if n == 1 && err == nil {
				break
			}
		}
		conv(t)
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			conv(t)
		}
	}
}
