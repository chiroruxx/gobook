package practice11_1

import (
	"testing"
	"unicode/utf8"
)

type utfLenType = [utf8.UTFMax + 1]int

func Test_charCount(t *testing.T) {
	tests := []struct {
		name         string
		input        []byte
		wantCounts   map[rune]int
		wantUTFLen   utfLenType
		wantInvalids int
	}{
		{
			"default",
			[]byte("abc"),
			map[rune]int{
				'a': 1,
				'b': 1,
				'c': 1,
			},
			[utf8.UTFMax + 1]int{
				0, 3, 0, 0, 0,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCounts, gotUTFLen, gotInvalids := charCount(tt.input)
			if !assertEqualsMap(tt.wantCounts, gotCounts) {
				t.Errorf("charCount() != %v, got %v", tt.wantCounts, gotCounts)
			}
			if tt.wantUTFLen != gotUTFLen {
				t.Errorf("charCount() != %v, got %v", tt.wantUTFLen, gotUTFLen)
			}
			if tt.wantInvalids != gotInvalids {
				t.Errorf("charCount() != %v, got %v", tt.wantInvalids, gotInvalids)
			}
		})
	}
}

func assertEqualsMap[K, V comparable](a, b map[K]V) bool {
	if len(a) != len(b) {
		return false
	}

	for key, value := range a {
		found, ok := b[key]
		if !ok || value != found {
			return false
		}
	}

	return true
}
