package substringconcatwords

import "testing"

func TestFindMatchingSubstrings(t *testing.T) {

	// s := "barfoothefoobarman"
	s := "wordgoodgoodgoodbestword"
	// words := []string{"foo", "bar"}
	words := []string{"word", "good", "best", "word"}

	resultPos := findMatchingSubstrings(s, words)
	t.Logf("found matching in positions = %#v", resultPos)
}
