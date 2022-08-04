package sexpr

import (
	"strings"
	"testing"
)

// TestFloat tests floats encoding as S-Expression notation
func TestFloat(t *testing.T) {
	type Test struct {
		value    interface{}
		expected string
	}
	tests := []Test{
		{float32(3), "3"},
		{float32(3.4), "3.4"},
		{float32(0), "0"},
		{float32(0.4), "0.4"},
		{float32(-10.1), "-10.1"},
		{float64(4), "4"},
		{float64(4.1), "4.1"},
		{float64(0), "0"},
		{float64(0.3), "0.3"},
		{float64(-12.03), "-12.03"},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if strings.Compare(string(data), test.expected) != 0 {
			t.Fatalf("Bad marshaling: expected %q, obtained %q", test.expected, string(data))
		}
		t.Logf("Marshal() = %s\n", data)
	}
}

// TestComplex tests complexes encoding as S-Expression notation
func TestComplex(t *testing.T) {

	type Test struct {
		value    interface{}
		expected string
	}
	tests := []Test{
		{complex64(3), "#C(3 0)"},
		{complex64(3.4), "#C(3.4 0)"},
		{complex64(3i), "#C(0 3)"},
		{complex64(3.4i), "#C(0 3.4)"},
		{complex64(3 + 4i), "#C(3 4)"},
		{complex64(-3 - 4i), "#C(-3 -4)"},
		{complex64(3.2 - 4.1i), "#C(3.2 -4.1)"},
		{complex128(3), "#C(3 0)"},
		{complex128(3.4), "#C(3.4 0)"},
		{complex128(3i), "#C(0 3)"},
		{complex128(3.4i), "#C(0 3.4)"},
		{complex128(3 + 4i), "#C(3 4)"},
		{complex128(-3 - 4i), "#C(-3 -4)"},
		{complex128(3.2 - 4.1i), "#C(3.2 -4.1)"},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if strings.Compare(string(data), test.expected) != 0 {
			t.Fatalf("Bad marshaling: expected %q, obtained %q", test.expected, string(data))
		}
		t.Logf("Marshal() = %s\n", data)
	}
}

// TestBool tests booleans encoding as S-Expression notation
func TestBool(t *testing.T) {
	type Test struct {
		value    interface{}
		expected string
	}
	tests := []Test{
		{true, "t"},
		{false, "nil"},
	}

	for _, test := range tests {
		// Encode it
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if strings.Compare(string(data), test.expected) != 0 {
			t.Fatalf("Bad marshaling: expected %q, obtained %q", test.expected, string(data))
		}
		t.Logf("Marshal() = %s\n", data)
	}
}

// TestInterface tests interfaces encoding as S-Expression notation
func TestInterface(t *testing.T) {

	type Sample struct {
		dummy int
		str   string
	}
	type Test struct {
		value    interface{}
		expected string
	}
	tests := []Test{
		{[]int{1, 2, 3}, "((i (\"[]int\" (1 2 3))))"},
		{Sample{10, "MyString"}, "((i (\"sexpr.Sample\" ((dummy 10) (str \"MyString\")))))"},
	}

	for _, test := range tests {
		type Struct struct {
			i interface{}
		}
		s := Struct{test.value}
		// Encode it
		data, err := Marshal(s)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if strings.Compare(string(data), test.expected) != 0 {
			t.Fatalf("Bad marshaling: expected %q, obtained %q", test.expected, string(data))
		}
		t.Logf("Marshal() = %s\n", data)
	}
}
