package khighest

import "testing"

func TestFindKHighest(t *testing.T) {

	t.Logf("Kth highest number=%d", findKHighest(3))

	t.Logf("After adding 6, Kth highest number=%d", Add(6))
	t.Logf("After adding 13, Kth highest number=%d", Add(13))
}
