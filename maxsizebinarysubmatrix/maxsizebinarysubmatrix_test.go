package maxsizebinarysubmatrix

import "testing"

func TestFindLargestSubMatrix(t *testing.T) {
	x, y, count := findLargestSubMatrix()
	t.Logf("Max count found at (%d,%d) = %d", x, y, count)
}
