// intset manipulate a set of integers as a bit array
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
	fmt.Printf("Len(x)=%d\n", x.Len())                               //"4"
	x.Remove(10)
	x.Remove(9)
	x.Remove(192)
	fmt.Printf("x=%s\n", x.String()) // "{1 42 144}"
	x.Clear()
	fmt.Printf("x=%s\n", x.String()) // "{}"
	z := y.Copy()
	fmt.Printf("z=%s\n", z.String()) // "{9 42}"
}
