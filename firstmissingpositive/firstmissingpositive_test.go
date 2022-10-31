package firstmissingpositive

import "testing"

func TestFindFirstMissingPositive(t *testing.T) {

	// input := []int{1, 2, 4, 0}
	input := []int{3, 4, -1, 1}
	// input := []int{7, 8, 9, 11, 12}
	t.Logf("first missing positive = %d", findFirstMissingPositive(input))
}
