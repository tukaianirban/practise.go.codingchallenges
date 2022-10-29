package stringsegment

import "testing"

func TestFindStringSegment(t *testing.T) {

	var dictionary = [4]string{"apple", "pear", "pier", "pie"}
	var masterWord = "applepie"
	word1, word2 := findStringSegment(masterWord, dictionary[:])

	t.Logf("found words= %s %s", word1, word2)
}

func TestIsContainedInStringArray(t *testing.T) {

	var dictionary = [6]string{"apple", "pear", "pier", "pie", "fruit", "banana"}
	var findWord = "per"
	var excepts = [3]string{"pear"}

	t.Logf("isContainedInStringArray = %t", isContainedInStringArray(dictionary[:], findWord, excepts[:]))
}
