// Package equal provides a deep equivalence relation for arbitrary values.
package equal

import (
	"math"
	"reflect"
	"unsafe"
)

// comparison is used to store comparison history in order to prevent infinite loops
type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

// cmp check if x and y values are near
func cmp(x, y float64) bool {
	if x == y {
		return true
	}
	delta := math.Abs(x-y) * 1000000000
	if math.Abs(x) < math.Abs(y) {
		return delta < math.Abs(y)
	} else {
		return delta < math.Abs(x)
	}
}

// equal checks if two values are deeply equal
// Thanks to the comparison map, infinite loops are prevented
func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	// Check values validity
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}

	// Check if values have the same type
	if x.Type() != y.Type() {
		return false
	}

	// Cycle check to prevent infinite loops
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // identical references
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	// Compare values according to their kind
	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return cmp(float64(x.Int()), float64(y.Int()))

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return cmp(float64(x.Uint()), float64(y.Uint()))

	case reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Float32, reflect.Float64:
		return cmp(x.Float(), y.Float())

	case reflect.Complex64, reflect.Complex128:
		realEqual := cmp(real(x.Complex()), real(y.Complex()))
		imagEqual := cmp(imag(x.Complex()), imag(y.Complex()))
		return realEqual && imagEqual

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, k := range x.MapKeys() {
			if !equal(x.MapIndex(k), y.MapIndex(k), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

// Equal reports whether x and y are deeply equal.
//
// Map keys are always compared with ==, not deeply.
// (This matters for keys containing pointers or interfaces.)
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}
