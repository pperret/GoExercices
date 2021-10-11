// counter is a program counting words and lines
package main

import (
	"bufio"
	"fmt"
)

// Counters is the internal structure used to count words and lines
// Due to line breaks and word breaks that can be anywhere, it is not possible
// to count lines and words on the fly. So, a buffer is used.
type Counters struct {
	buffer []byte
}

// Write appends bytes to the Counters's buffer
// Therefore, Counters implements the io.Writer standard interface
func (c *Counters) Write(p []byte) (int, error) {
	c.buffer = append(c.buffer, p...)
	return len(p), nil
}

// Count is a generic function to count anything in the Counters's buffer
func (c *Counters) Count(fct func([]byte, bool) (int, []byte, error)) (int, error) {
	pos := 0
	count := 0
	for {
		advance, token, err := fct(c.buffer[pos:], true)
		if err != nil {
			return 0, err
		}
		if token == nil {
			break
		}
		pos += advance
		count++
	}
	return count, nil
}

// Words returns the number of words in the Counters's buffer
func (c *Counters) Words() (int, error) {
	return c.Count(bufio.ScanWords)
}

// Lines returns the number of lines in the Counters's buffer
func (c *Counters) Lines() (int, error) {
	return c.Count(bufio.ScanLines)
}

// main is the entry point of the program
func main() {
	var counter Counters
	fmt.Fprintf(&counter, "This is a sample\n")
	fmt.Fprintf(&counter, "program to ")
	fmt.Fprintf(&counter, "count wo")
	fmt.Fprintf(&counter, "rds and lines\n Is not it?")

	words, _ := counter.Words()
	lines, _ := counter.Lines()
	fmt.Println(words, lines) // Should be "13 3"
}
