// Xmlselect prints the text of selected elements of an XML document.
package sexpr

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDecoder(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	// Encode it
	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	// Create the reader
	reader := bytes.NewReader(data)

	dec := NewDecoder(reader)
	var level int // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			if level != 0 {
				t.Fatalf("Lists are not closed (level=%d)\n", level)
			}
			break
		} else if err != nil {
			t.Fatalf("get token: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case StartList:
			t.Logf("%*sStartList\n", level*2, "")
			level++
		case EndList:
			if level <= 0 {
				t.Fatalf("No list is open (level=%d)\n", level)
			}
			level--
			t.Logf("%*sEndList\n", level*2, "")
		case Symbol:
			t.Logf("%*sSymbol(%q)\n", level*2, "", tok)
		case String:
			t.Logf("%*sString(%q)\n", level*2, "", tok)
		case Int:
			t.Logf("%*sInt(%d)\n", level*2, "", tok)
		}
	}
}
