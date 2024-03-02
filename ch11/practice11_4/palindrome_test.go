package practice11_4

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindromeWithSymbol(rng *rand.Rand) string {
	symbols := []rune{',', ' ', '.'}
	n := rand.Intn(25)

	runes := make([]rune, n)
	for i, j := 0, n-1; i < j; {
		insertSymbol := rng.Intn(2)
		if insertSymbol == 0 {
			r := symbols[rng.Intn(len(symbols))]
			runes[i] = r
			i++
		}

		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[j] = r
		i++
		j--

		insertSymbol = rng.Intn(2)
		if insertSymbol == 0 {
			r := symbols[rng.Intn(len(symbols))]
			runes[j] = r
			j--
		}
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindromeWithSymbol(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) == false", p)
		}
	}
}
