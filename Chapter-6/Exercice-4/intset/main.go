// intset manipulates a set of integers as a bit array
package main

import "fmt"

// main is the entry point of the program
func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Printf("x=%s\n", x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Printf("y=%s\n", y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Printf("x=%s\n", x.String())                                 // "{1 9 42 144}"
	fmt.Printf("x.Has(9)=%t, x.Has(123)=%t\n", x.Has(9), x.Has(123)) // "true false"
	elems := x.Elems()
	for _, e := range elems {
		fmt.Printf("%d ", e)
	}
	fmt.Println() // 1 9 42 144
}
