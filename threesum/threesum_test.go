package threesum

import "testing"

func TestFind3Sums(t *testing.T) {

	inputs := [][]int{
		{-1, 0, 1, 2, -1, -4},
	}

	for _, input := range inputs {
		t.Logf("Input=%#v  3Sum result=%#v", input, find3Sums(input))
	}
}
