// panic is a sample that implements a function returning a value without return instruction
package main

import "fmt"

// main is the entry point of the program
func main() {

	fmt.Printf("test(): %d\n", test())
}

// test is the function returning a value without return instruction
func test() (val int) {
	defer func() {
		v := recover()
		if v != nil {
			v1, ok := v.(int)
			if !ok {
				panic(v)
			}
			val = v1
		}
	}()

	panic(12)
}
