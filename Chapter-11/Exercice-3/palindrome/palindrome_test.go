package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

// Thanks to Torbiak, palidrome and non-palidrome generation is based on Eli Bendersky's excellent article
// See http://eli.thegreenplace.net/2010/01/28/generating-random-sentences-from-a-context-free-grammar/
// NON = acb | ab | aNONa | aNONb | aPALb
// PAL = eps | a | aa | aPALa
// However, my implemention is different of the Torbiak's one

// randomLetter generates a random letter as a rune using the pseudo-random number generator.
func randomLetter(rng *rand.Rand) rune {
	for {
		r := rune(rng.Intn(0x1000))
		if unicode.IsLetter(r) {
			return r
		}
	}
}

// randomTwoLetters generates two different random letters as runes using
// the pseudo-random number generator.
func randomTwoLetters(rng *rand.Rand) (rune, rune) {
	a := randomLetter(rng)
	al := unicode.ToLower(a)
	for {
		b := randomLetter(rng)
		if al != unicode.ToLower(b) {
			return a, b
		}
	}
}

// randomPalindrome returns a palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	switch rng.Intn(20) {
	case 0: // eps
		return ""
	case 1: // a
		return string(randomLetter(rng))
	case 2: // aa
		a := randomLetter(rng)
		return string(a) + string(a)
	default: // aPALa
		a := randomLetter(rng)
		return string(a) + randomPalindrome(rng) + string(a)
	}
}

// randomNonPalindrome returns a non-palindrome whose length and contents
// are derived from the pseudo-random number generator rng.
func randomNonPalindrome(rng *rand.Rand) string {
	switch rng.Intn(10) {
	case 0: // ab
		a, b := randomTwoLetters(rng)
		return string(a) + string(b)
	case 1, 2, 3: // aNONa
		a := randomLetter(rng)
		return string(a) + randomNonPalindrome(rng) + string(a)
	case 4, 5, 6: // aNONb
		a, b := randomTwoLetters(rng)
		return string(a) + randomNonPalindrome(rng) + string(b)
	default: // aPALb
		a, b := randomTwoLetters(rng)
		return string(a) + randomPalindrome(rng) + string(b)
	}
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = true", p)
		}
	}
}
