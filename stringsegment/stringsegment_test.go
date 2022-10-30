package stringsegment

import "testing"

func TestFindSegments(t *testing.T) {

	var dictionary = []string{"apple", "pear", "pier", "pie", "orange", "mango"}
	var masterWord = "pierapple"

	segments := findSegments(masterWord, dictionary[:])
	t.Logf("Found segments = %#v", segments)
}

func TestFindSegments2(t *testing.T) {

	var dictionary = []string{"apple", "pear", "pier", "pie", "orange", "mango"}
	var masterWord = "pierapporange"

	segments := findSegments2(masterWord, dictionary)
	t.Logf("Found segments = %#v", segments)
}
