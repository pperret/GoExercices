package sexpr

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

// A Token includes Symbol, String, Int, StartList and EndList
type Token interface{}
type Symbol string
type Int int64
type String string
type StartList struct{}
type EndList struct{}

// Decoder is the decoder data
type Decoder struct { /* ... */
	scan  scanner.Scanner
	token rune // the current token
}

// NewDecoder creates a new decoder instance
func NewDecoder(reader io.Reader) *Decoder {
	dec := Decoder{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec.scan.Init(reader)
	return &dec
}

// Token returns next Token in sequence
func (dec *Decoder) Token() (Token, error) {
	// Get the next token
	dec.token = dec.scan.Scan()

	// Manage the token according its type
	switch dec.token {
	case scanner.Ident:
		return Symbol(dec.scan.TokenText()), nil
	case scanner.String:
		s, err := strconv.Unquote(dec.scan.TokenText())
		if err != nil {
			return nil, err
		}
		return String(s), nil
	case scanner.Int:
		i, err := strconv.Atoi(dec.scan.TokenText())
		if err != nil {
			return nil, err
		}
		return Int(int64(i)), nil
	case '(':
		return StartList{}, nil
	case ')':
		return EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected token %q", dec.scan.TokenText())
	}
}
