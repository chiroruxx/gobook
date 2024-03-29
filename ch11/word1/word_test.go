package word1

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("detarated") {
		t.Error(`IsPalindrome("detarated") == false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") == false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") == true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") == false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome("%q") == false`, input)
	}
}
