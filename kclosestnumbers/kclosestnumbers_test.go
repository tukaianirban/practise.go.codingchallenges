package kclosestnumbers

import "testing"

func TestFindIndexes(t *testing.T) {

	var numbers = [7]int{3, 5, 7, 8, 9, 11, 14}
	b, e, err := findIndexes(numbers[:], 0, len(numbers)-1, 15)
	t.Logf("begin=%d end=%d err=%s", b, e, err)
}

func TestFindClosest(t *testing.T) {

	var numbers = [7]int{3, 6, 7, 8, 11, 12, 14}
	closes := findClosest(numbers[:], 3, 4)
	t.Logf("closest numbers to 10 are = %#v", closes)
}
