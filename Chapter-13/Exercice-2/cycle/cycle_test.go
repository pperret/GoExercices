package cycle

import (
	"testing"
)

func TestCycle(t *testing.T) {
	one := 1
	ch1 := make(chan int)
	var iface1 interface{} = &one
	type CycleSlice []CycleSlice
	var cycleSlice = make(CycleSlice, 1)
	cycleSlice[0] = cycleSlice
	type CyclePtr *CyclePtr
	var cyclePtr1 CyclePtr
	cyclePtr1 = &cyclePtr1

	for _, test := range []struct {
		x    interface{}
		want bool
	}{
		// int
		{0, false},
		// float
		{0.0, false},
		// complex
		{complex(0, 0), false},
		// string
		{"foo", false},
		// slices
		{[]string{"foo"}, false},
		// maps
		{map[string][]int{"foo": {1, 2, 3}}, false},
		// pointers
		{&one, false},
		// functions
		{(func())(nil), false},
		// arrays
		{[...]int{1, 2, 3}, false},
		// channels
		{ch1, false},
		// interfaces
		{&iface1, false},
		// structs
		{struct {
			a int
			b string
		}{1, "foo"}, false},
		// slice cycles
		{cycleSlice, true},
		// pointer cycles
		{cyclePtr1, true},
	} {
		if Cycle(test.x) != test.want {
			t.Errorf("Cycle(%v) = %t",
				test.x, !test.want)
		}
	}
}
