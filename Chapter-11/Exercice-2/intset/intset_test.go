package main

import (
	"strings"
	"testing"
)

func TestHas(t *testing.T) {
	var x IntSet
	var y IntMap = make(IntMap)

	if x.Has(1) != y.Has(1) {
		t.Errorf("IntSet and IntMap disagree on having 1")
	}

	x.Add(2)
	y.Add(2)

	if x.Has(2) != y.Has(2) {
		t.Errorf("IntSet and IntMap disagree on having 2")
	}

	x.Add(3)
	if x.Has(3) == y.Has(3) {
		t.Errorf("IntSet and IntMap agree on having 3")
	}

	y.Add(4)
	if x.Has(4) == y.Has(4) {
		t.Errorf("IntSet and IntMap agree on having 4")
	}
}

func TestAdd(t *testing.T) {
	var x IntSet
	var y IntMap = make(IntMap)

	x.Add(2)
	y.Add(2)

	x.Add(10)
	y.Add(10)

	x.Add(1000)
	y.Add(1000)

	xs := x.String()
	ys := y.String()
	if strings.Compare(xs, ys) != 0 {
		t.Errorf("Additions in IntSet and IntMap differ: '%s' / '%s'", xs, ys)
	}
}

func TestUnion(t *testing.T) {
	var x1 IntSet
	var x2 IntSet
	var y1 IntMap = make(IntMap)
	var y2 IntMap = make(IntMap)

	x1.Add(10)
	y1.Add(10)

	x1.Add(20)
	y1.Add(20)

	x1.Add(30)
	y1.Add(30)

	x2.Add(15)
	y2.Add(15)

	x2.Add(20)
	y2.Add(20)

	x2.Add(25)
	y2.Add(25)

	x1.UnionWith(&x2)
	y1.UnionWith(&y2)

	xs := x1.String()
	ys := y1.String()
	if strings.Compare(xs, ys) != 0 {
		t.Errorf("Unions in IntSet and IntMap differ: '%s' / '%s'", xs, ys)
	}
}
