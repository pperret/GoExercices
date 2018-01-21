// Package echo contains several implementations of echo program
package echo

import (
	"os"
	"strings"
)

// Method1 is an un-efficient version
func Method1() string {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

// Method2 is an efficient version
func Method2() string {
	return strings.Join(os.Args[1:], " ")
}
