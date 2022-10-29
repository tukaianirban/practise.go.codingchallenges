package longestsubstrpalindrome

import "testing"

func TestFindLongestPalindromeSubstring(t *testing.T) {

	strs := []string{"babad", "cbbd", "abacdfgdcaba"}
	for _, str := range strs {
		// t.Logf("longest palindrome in %s is %s", str, findLongestPalindromeSubstring(str))
		t.Logf("longest palindrome in %s is %s", str, findLongestPalindromeSubstring2(str))
	}
}
