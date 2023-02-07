package equal

import (
	"testing"
)

func TestEqual(t *testing.T) {

	for _, test := range []struct {
		x, y interface{}
		want bool
	}{
		// int
		{0, 0, true},
		{1, 1, true},
		{1, 2, false},                     // different values
		{1000000000, 1000000002, false},   // more than 1/1000000000
		{2000000000, 2000000001, true},    // less than 1/1000000000
		{-1000000000, -1000000002, false}, // more than 1/1000000000
		{-2000000000, -2000000001, true},  // less than 1/1000000000
		// float
		{0.0, 0.0, true},
		{1.0, 1.0, true},
		{1.0, 2.0, false},           // different values
		{1.0, 1.000000002, false},   // more than 1/1000000000
		{2.0, 2.000000001, true},    // less than 1/1000000000
		{-1.0, -1.000000002, false}, // more than 1/1000000000
		{-2.0, -2.000000001, true},  // less than 1/1000000000
		// complex
		{complex(0, 0), complex(0, 0), true},
		{complex(1, 0), complex(1, 0), true},
		{complex(1, 0), complex(0, 1), false},
		{complex(1, 0), complex(1.000000002, 0), false},
		{complex(2, 0), complex(2.000000001, 0), true},
		{complex(-1, 0), complex(-1.000000002, 0), false},
		{complex(-2, 0), complex(-2.000000001, 0), true},
		{complex(0, 1), complex(0, 1.000000002), false},
		{complex(0, 2), complex(0, 2.000000001), true},
		{complex(0, -1), complex(0, -1.000000002), false},
		{complex(0, -2), complex(0, -2.000000001), true},
	} {
		if Equal(test.x, test.y) != test.want {
			t.Errorf("Equal(%v, %v) = %t",
				test.x, test.y, !test.want)
		}
	}
}
