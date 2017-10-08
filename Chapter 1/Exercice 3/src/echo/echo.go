// Several implementations of echo program
package echo

import (
	"os"
	"strings"
)

// Un-efficient version
func Method1() (string){
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// Efficient version
func Method2() (string) {
	return strings.Join(os.Args[1:], " ")
}
