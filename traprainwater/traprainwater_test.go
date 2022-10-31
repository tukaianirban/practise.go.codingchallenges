package traprainwater

import "testing"

func TestGetTrappedRainWater(t *testing.T) {

	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	t.Logf("total volume of rain water captured = %d", getTrappedRainWater(input))
}
