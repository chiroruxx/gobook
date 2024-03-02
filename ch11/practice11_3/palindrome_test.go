package practice11_3

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rand.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rand.Intn(23) + 2
	runes := make([]rune, n)

	var first, last *rune
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))

		if unicode.IsLetter(r) {
			last = &r
			if first == nil {
				first = &r
			}
		}

		runes[i] = r
	}

	if first == nil || last == nil {
		return randomNonPalindrome(rng)
	}

	if unicode.ToLower(*first) == unicode.ToLower(*last) {
		return randomNonPalindrome(rng)
	}

	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) == false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) == true, at %d, len %d", p, i, len(p))
		}
	}
}
