package main

import (
	"flag"
	"math/rand"
	"testing"
)

// Test parameters used to analyze API behavior based on range of values and count of items
var max_range = flag.Int("max", 1000000, "Max value in the set")
var count = flag.Int("count", 1000, "Count of integers in the set")

// getValue returns a value that is not already in the set of integers
func getValue(intSet64 *IntSet64, max int) int {
	for {
		v := rand.Intn(max)
		if !intSet64.Has(v) {
			return v
		}
	}
}

// generateSets generates the same set of integers as IntSet64, IntSet32 and IntMap
func generateSets(max, count int) (*IntSet64, *IntSet32, *IntMap) {
	var intMap = make(IntMap)
	var intSet64 IntSet64
	var intSet32 IntSet32
	for i := 0; i < count; i++ {
		for {
			v := rand.Intn(max)
			if !intSet64.Has(v) {
				intSet64.Add(v)
				intSet32.Add(v)
				intMap.Add(v)
				break
			}
		}
	}
	return &intSet64, &intSet32, &intMap
}

// Data for Dup benchmarks
// Same data are used for IntSet64, IntSet32 and IntMap implementations
var dupIntSet64Data *IntSet64
var dupIntSet32Data *IntSet32
var dupIntMapData *IntMap

// Initialize data for Dup benchmarks
func dupInit(b *testing.B) {
	if dupIntSet64Data == nil {
		dupIntSet64Data, dupIntSet32Data, dupIntMapData = generateSets(*max_range, *count)
		b.ResetTimer()
	}
}

// BenchmarkDupSet64 benchs the Dup function of the IntSet64 implementation
func BenchmarkDupSet64(b *testing.B) {
	dupInit(b)
	for n := 0; n < b.N; n++ {
		dupIntSet64Data.Dup()
	}
}

// BenchmarkDupSet32 benchs the Dup function of the IntSet32 implementation
func BenchmarkDupSet32(b *testing.B) {
	dupInit(b)
	for n := 0; n < b.N; n++ {
		dupIntSet32Data.Dup()
	}
}

// BenchmarkDupMap benchs the Dup function of the IntMap implementation
func BenchmarkDupMap(b *testing.B) {
	dupInit(b)
	for n := 0; n < b.N; n++ {
		dupIntMapData.Dup()
	}
}

// Data for Has benchmarks
// Same data are used for IntSet64, IntSet32 and IntMap implementations
var hasIntSet64Data *IntSet64
var hasIntSet32Data *IntSet32
var hasIntMapData *IntMap

// Initialize data for Has benckmarks
func hasInit(b *testing.B) {
	if hasIntSet64Data == nil {
		hasIntSet64Data, hasIntSet32Data, hasIntMapData = generateSets(*max_range, *count)
		b.ResetTimer()
	}
}

// BenchmarkHasSet64 benchs the Has function of the IntSet64 implementation
func BenchmarkHasSet64(b *testing.B) {
	hasInit(b)
	for n := 0; n < b.N; n++ {
		hasIntSet64Data.Has(1000)
	}
}

// BenchmarkHasSet32 benchs the Has function of the IntSet32 implementation
func BenchmarkHasSet32(b *testing.B) {
	hasInit(b)
	for n := 0; n < b.N; n++ {
		hasIntSet32Data.Has(1000)
	}
}

// BenchmarkHasMap benchs the Has function of the IntMap implementation
func BenchmarkHasMap(b *testing.B) {
	hasInit(b)
	for n := 0; n < b.N; n++ {
		hasIntMapData.Has(1000)
	}
}

// Data for Add benchmarks
// Same data are used for IntSet64, IntSet32 and IntMap implementations
var addIntSet64Data *IntSet64
var addIntSet32Data *IntSet32
var addIntMapData *IntMap
var addValue int

// Initialize data for Add benchmarks
func addInit(b *testing.B) {
	if addIntSet64Data == nil {
		addIntSet64Data, addIntSet32Data, addIntMapData = generateSets(*max_range, *count)
		addValue = getValue(addIntSet64Data, *max_range)
		b.ResetTimer()
	}
}

// BenchmarkAddSet64 benchs the Add function of the IntSet64 implementation
func BenchmarkAddSet64(b *testing.B) {
	addInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := addIntSet64Data.Dup()
		b.StartTimer()
		data.Add(addValue)
	}
}

// BenchmarkAddSet32 benchs the Add function of the IntSet32 implementation
func BenchmarkAddSet32(b *testing.B) {
	addInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := addIntSet32Data.Dup()
		b.StartTimer()
		data.Add(addValue)
	}
}

// BenchmarkAddMap benchs the Add function of the IntMap implementation
func BenchmarkAddMap(b *testing.B) {
	addInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := addIntMapData.Dup()
		b.StartTimer()
		data.Add(addValue)
	}
}

// Data for UnionWith benchmarks
// Same data are used for IntSet64, IntSet32 and IntMap implementations
var unionIntSet64Data1, unionIntSet64Data2 *IntSet64
var unionIntSet32Data1, unionIntSet32Data2 *IntSet32
var unionIntMapData1, unionIntMapData2 *IntMap

// Initialize data for UnionWith benchmarks
func unionWithInit(b *testing.B) {
	if unionIntSet64Data1 == nil {
		unionIntSet64Data1, unionIntSet32Data1, unionIntMapData1 = generateSets(*max_range, *count)
		unionIntSet64Data2, unionIntSet32Data2, unionIntMapData2 = generateSets(*max_range, *count)
		b.ResetTimer()
	}
}

// BenchmarkUnionSet64 benchs the Union function of the IntSet64 implementation
func BenchmarkUnionSet64(b *testing.B) {
	unionWithInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := unionIntSet64Data1.Dup()
		b.StartTimer()
		data.UnionWith(unionIntSet64Data2)
	}
}

// BenchmarkUnionSet32 benchs the Union function of the IntSet32 implementation
func BenchmarkUnionSet32(b *testing.B) {
	unionWithInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := unionIntSet32Data1.Dup()
		b.StartTimer()
		data.UnionWith(unionIntSet32Data2)
	}
}

// BenchmarkUnionMap benchs the Union function of the IntMap implementation
func BenchmarkUnionMap(b *testing.B) {
	unionWithInit(b)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		data := unionIntMapData1.Dup()
		b.StartTimer()
		data.UnionWith(unionIntMapData2)
	}
}
