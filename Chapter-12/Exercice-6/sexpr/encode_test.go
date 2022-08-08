package sexpr

import (
	"testing"
)

// The test does not make direct assertions about the encoded output
// because the output depends on map iteration order, which is
// nondeterministic.  The output of the t.Log statements can be
// inspected by running the test with the -v flag:
func Test(t *testing.T) {
	type Movie struct {
		str         string
		emptyStr    string
		integer     int
		nullInteger int
		nullPtr     *string
		stru        struct {
			nullInteger int
			emptyStr    string
		}
	}
	strangelove := Movie{
		str:         "non-empty string",
		emptyStr:    "",
		integer:     123,
		nullInteger: 0,
		nullPtr:     nil,
		stru: struct {
			nullInteger int
			emptyStr    string
		}{
			nullInteger: 0,
			emptyStr:    "",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)
}
