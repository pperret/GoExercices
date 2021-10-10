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
	x.IntersectWith(&y)
	fmt.Printf("x=%s\n", x.String()) // "{9 42}"
	x.Add(1000)
	x.Add(3)
	x.DifferenceWith(&y)
	fmt.Printf("x=%s\n", x.String()) // "{3 1000}"
	x.SymmetricDifference(&y)
	fmt.Printf("x=%s\n", x.String()) // "{3 9 42 1000}"
}
