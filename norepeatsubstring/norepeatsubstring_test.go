package norepeatsubstring

import "testing"

func TestFindLongestSubstring(t *testing.T) {

	str := "abcabcbb"
	t.Logf("longest substring matching in %s is = %d", str, lengthOfLongestSubstring(str))

	str = "bbbbb"
	t.Logf("longest substring matching in %s is = %d", str, lengthOfLongestSubstring(str))

	str = "pwwkew"
	t.Logf("longest substring matching in %s is = %d", str, lengthOfLongestSubstring(str))
}
