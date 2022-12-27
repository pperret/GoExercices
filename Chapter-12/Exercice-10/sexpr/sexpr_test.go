package sexpr

import (
	"reflect"
	"testing"
)

// TestFloat tests floats decoding as S-Expression notation
func TestFloat(t *testing.T) {
	type Test struct {
		value    interface{}
		expected float64
	}
	tests := []Test{
		{float32(3), 3},
		{float32(3.4), 3.4},
		{float32(0), 0},
		{float32(0.4), 0.4},
		{float32(-10.1), -10.1},
		{float64(4), 4},
		{float64(4.1), 4.1},
		{float64(0), 0},
		{float64(0.3), 0.3},
		{float64(-12.03), -12.03},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		// Decode it
		var value float64
		err = Unmarshal(data, &value)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		// Check equality.
		if !reflect.DeepEqual(value, test.expected) {
			t.Fatalf("not equal %v/%v", value, test.expected)
		}
		t.Logf("Unmarshal() = %+v\n", value)
	}
}

// TestComplex tests complexes decoding as S-Expression notation
func TestComplex(t *testing.T) {

	type Test struct {
		value    interface{}
		expected complex128
	}
	tests := []Test{
		{complex64(3), 3},
		{complex64(3.4), 3.4},
		{complex64(3i), 3i},
		{complex64(3.4i), 3.4i},
		{complex64(3 + 4i), 3 + 4i},
		{complex64(-3 - 4i), -3 - 4i},
		{complex64(3.2 - 4.1i), 3.2 - 4.1i},
		{complex128(3), 3},
		{complex128(3.4), 3.4},
		{complex128(3i), 3i},
		{complex128(3.4i), 3.4i},
		{complex128(3 + 4i), 3 + 4i},
		{complex128(-3 - 4i), -3 - 4i},
		{complex128(3.2 - 4.1i), 3.2 - 4.1i},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		// Decode it
		var value complex128
		err = Unmarshal(data, &value)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		// Check equality.
		if !reflect.DeepEqual(value, test.expected) {
			t.Fatalf("not equal: %v/%v", value, test.expected)
		}
		t.Logf("Unmarshal() = %+v\n", value)
	}
}

// TestBool tests booleans decoding as S-Expression notation
func TestBool(t *testing.T) {
	type Test struct {
		value interface{}
	}
	tests := []Test{
		{true},
		{false},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		// Decode it
		var value bool
		err = Unmarshal(data, &value)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		// Check equality.
		if !reflect.DeepEqual(value, test.value) {
			t.Fatalf("not equal: %v/%v", value, test.value)
		}
		t.Logf("Unmarshal() = %+v\n", value)
	}
}

// TestInterface tests interfaces encoding as S-Expression notation
func TestInterface(t *testing.T) {
	type Sample struct {
		Dummy int64
		Str   string
	}

	type Test struct {
		value interface{}
	}
	tests := []Test{
		{[]int{1, 2, 3}},
		{Sample{10, "MyString"}},
	}

	InterfaceMap["[]int"] = reflect.TypeOf([]int{})
	InterfaceMap["sexpr.Sample"] = reflect.TypeOf(Sample{0, ""})

	for _, test := range tests {
		type Struct struct {
			I interface{}
		}
		s := Struct{test.value}
		// Encode it
		data, err := Marshal(s)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		// Decode it
		var value Struct
		err = Unmarshal(data, &value)
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		// Check equality.
		if !reflect.DeepEqual(value, s) {
			t.Fatalf("not equal:  %v/%v", value, s)
		}
		t.Logf("Unmarshal() = %+v\n", value)
	}
}
