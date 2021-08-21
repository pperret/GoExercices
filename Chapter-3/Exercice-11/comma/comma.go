// Comma prints its argument numbers with a comma at each power of 1000.
//
//	$ ./comma 1 12 123 1234 1234567890 -1 -123 -1234 -1.234 -123.456 -1234.567

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a decimal number string.
func comma(s string) string {
	var buf bytes.Buffer

	// Length of the integer part
	p := strings.IndexByte(s, '.')
	if p < 0 {
		p = len(s) // Full lenght of the string if there is no decimal part
	}

	// Start position of digits
	pos := 0
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		pos = 1
		p--
	}

	// Compute first length
	m := p % 3
	if m == 0 {
		m = 3
	}
	m += pos
	buf.WriteString(s[:m]) // Including the optional sign

	// Next parts separated by commas
	for ; m < p+pos; m += 3 {
		buf.WriteString(",")
		buf.WriteString(s[m : m+3])
	}

	// Append the decimal part (no comma)
	buf.WriteString(s[m:])

	return buf.String()
}
