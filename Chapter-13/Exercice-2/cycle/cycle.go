// Package cycle detects cycles in values
package cycle

import (
	"fmt"
	"reflect"
	"unsafe"
)

// comparison is used to store comparison history in order to detect cycles
type comparison struct {
	x unsafe.Pointer
	t reflect.Type
}

// cycle checks if a cycle is detected thanks to the comparison map
func cycle(x reflect.Value, seen map[comparison]bool) bool {
	// Check to detect cycles in the current value
	if x.CanAddr() {
		c := comparison{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if seen[c] {
			return true // already seen
		}
		seen[c] = true
	}

	// Checks cycles in values according to their kind
	switch x.Kind() {
	case reflect.Bool,
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32, reflect.Float64,
		reflect.Complex64, reflect.Complex128,
		reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return false

	case reflect.Ptr, reflect.Interface:
		return cycle(x.Elem(), seen)

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if cycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if cycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if cycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	default:
		panic(fmt.Sprintf("Unmanaged type of value %v", x.Kind()))
	}
}

// Cycle reports cycle in value
func Cycle(x interface{}) bool {
	seen := make(map[comparison]bool)
	return cycle(reflect.ValueOf(x), seen)
}
