package palindrome

import "testing"

func TestFindPalindromeSubstrings(t *testing.T) {

	results := findPalindromeSubstrings("aabbbaa")
	t.Logf("palindrome substrings = %#v", results)

}
