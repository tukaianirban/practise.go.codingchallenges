package mostwatercontainer

import "testing"

func TestGetLargestContainer(t *testing.T) {

	// heights := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	heights := []int{2, 3, 10, 5, 7, 8, 9}
	t.Logf("largest volume container is %d", getLargestContainer(heights))
}
