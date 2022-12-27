// Package sexpr provides a means for converting Go objects to and from S-expressions.
package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"text/scanner"
)

// InterfaceMap is the set of interfaces that can be decoded
var InterfaceMap map[string]reflect.Type

func init() {
	InterfaceMap = make(map[string]reflect.Type)
}

// Unmarshal parses S-expression data and populates the variable
// whose address is in the non-nil pointer out.
func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next() // get the first token
	defer func() {
		// NOTE: this is not an example of ideal error handling.
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

// lexer is a wrapper for Scanner standard type
type lexer struct {
	scan  scanner.Scanner
	token rune // the current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: Not an example of good error handling.
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

// The read function is a decoder for a small subset of well-formed
// S-expressions.  For brevity of our example, it takes many dubious
// shortcuts.
//
// The parser assumes
// - that the S-expression input is well-formed; it does no error checking.
// - that the S-expression input corresponds to the type of the variable.
// - that all numbers in the input are non-negative decimal integers.
// - that all keys in ((key value) ...) struct syntax are unquoted symbols.
// - that the input does not contain dotted lists such as (1 2 . 3).
// - that the input does not contain Lisp reader macros such 'x and #'x.
//
// The reflection logic assumes
//   - that v is always a variable of the appropriate type for the
//     S-expression value.  For example, v must not be a boolean,
//     interface, channel, or function, and if v is an array, the input
//     must have the correct number of elements.
//   - that v in the top-level call to read has the zero value of its
//     type and doesn't need clearing.
//   - that if v is a numeric variable, it is a signed integer.
func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// The only valid identifiers are "nil", "t" and struct field names.
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		case "t":
			if v.Kind() == reflect.Bool {
				v.SetBool(true)
				lex.next()
				return
			}
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // NOTE: ignoring errors
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
		if v.CanInt() {
			v.SetInt(int64(i))
			lex.next()
			return
		}
		if v.CanFloat() {
			v.SetFloat(float64(i))
			lex.next()
			return
		}
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(f)
		lex.next()
		return

	case '#':
		lex.consume('#')
		if lex.text() != "C" {
			panic(fmt.Sprintf("not a complex value: %q", lex.text()))
		}
		lex.consume(scanner.Ident)
		lex.consume('(')
		real := reflect.New(reflect.TypeOf(float64(0))).Elem()
		read(lex, real)
		img := reflect.New(reflect.TypeOf(float64(0))).Elem()
		read(lex, img)
		c128 := complex(real.Float(), img.Float())
		v.SetComplex(c128)
		lex.consume(')')
		return

	// In Go, minus is a distinct token
	case '-':
		lex.next()
		switch lex.token {
		case scanner.Int:
			i, _ := strconv.Atoi(lex.text()) // NOTE: ignoring errors
			if v.CanInt() {
				v.SetInt(int64(-i))
				lex.next()
				return
			}
			if v.CanFloat() {
				v.SetFloat(float64(-i))
				lex.next()
				return
			}
		case scanner.Float:
			f, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetFloat(-f)
			lex.next()
			return
		}

	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // consume ')'
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

// readlist decodes a list into a variable
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}

	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		if lex.token != scanner.String {
			panic(fmt.Sprintf("got token %q, want name", lex.text()))
		}
		interfaceName, _ := strconv.Unquote(lex.text()) // NOTE: Ignoring errors
		lex.consume(scanner.String)
		interfaceType, ok := InterfaceMap[interfaceName]
		if !ok {
			panic(fmt.Sprintf("unknown interface name %v", interfaceName))
		}
		w := reflect.New(interfaceType)
		read(lex, reflect.Indirect(w))
		v.Set(reflect.Indirect(w))

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

// endList detects the end of a list
func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
