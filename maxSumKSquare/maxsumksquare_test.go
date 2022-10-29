package maxsumksquare

import "testing"

func TestTraverseMatrix(t *testing.T) {

	result := traverseMatrix(3)

	// print results
	for i := 0; i < len(result); i++ {
		t.Logf("%#v", result[i])
	}
}
